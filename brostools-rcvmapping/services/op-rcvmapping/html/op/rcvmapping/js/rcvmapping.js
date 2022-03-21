import { DEFAULT_LOCALE } from "./config.js";
import { URL_API_JWT_VERIFY } from "./config.js";
import { URL_API_JWT_SESSION } from "./config.js";
import { URL_API_USER_LOCALE } from "./config.js";
import { URL_API_MST_ADDRSET } from "./config.js";
import { URL_API_RCVMAPPING } from "./config.js";
import { URL_API_RCVMAPPING_REMOVE } from "./config.js";

Vue.use(window.VuejsDialog.main.default,{
    html: true
})
Vue.use(VueI18n)
Vue.use(VueMeta)
Vue.component('paginate', VuejsPaginate)
var i18n = new VueI18n({
    locale: DEFAULT_LOCALE,
    messages: {
        jp:{
            "PAGE_TITLE" : "受注紐付け設定",
            "CHANGE_LOCALE" : "言語の切り替え",
            "LOGIN_SESSION_INVALID":"ログインセッションが無効です。ログイン画面に移動します。",
            "SERVER_ACCESS_PROBLEM":"サーバーの通信で問題が起こりました。再読み込みをします。この画面が続く場合はシステム管理者にお問い合わせください。",
            "CONFIRM_DELETE": "この受信した請求書を削除してもよろしいですか？",
            "DELETE_COMPLETE": "削除しました。",
            "CONFIRM_OK": "削除する",
            "EDIT":"編集",
            "CONFIRM_NG": "キャンセル",
            "OK": 'OK',
            "LOGIN_SESSION_VALID_TITLE":"ログインセッションが無効です",
            "SERVER_ACCESS_PROBLEM_TITLE": "サーバーの通信で問題が起こりました",
            "REGISTER":"設定",
            "REMOVE":"解除",
            "RECEIVED_ORDOER_NO":"受注No",
            "ESTIMATIONS_NO":"見積No",
            "SAVE":"保存",
            "CLIENT":"クライアント",
            "ADDRESS_SET":"宛先セット",
            "VALIDATION_ERROR_TITLE": "エラー",
            "VALIDATION_ERROR": "入力フォームに不備があります。",
            "AREA_MAS": "受注Noまたは見積No",
            "POST_SUCCESS_TITLE": "完了",
            "POST_SUCCESS": "受注紐付け設定が保存されました。",
            "CONFIRM_ADD_GROUP": "確認",
            "CONFIRM_ADD_GROUP_MESSAGE": "登録してもよろしいですか？",
            "POST_ERROR_STRING_EXIT":"対象受注が存在しません。",
            "POST_ERROR_STRING":"対象受注の請求書が発行されました。",
            "CONFIRM_REMOVE_GROUP": "確認",
            "CONFIRM_REMOVE_GROUP_MESSAGE": "解除してもよろしいですか？",
            "REMOVE_SUCCESS_TITLE": "完了",
            "REMOVE_SUCCESS": "受注紐付け設定が解除されました。",
        },
        en:{
            "PAGE_TITLE" : "Received Order Address Set Setting",
            "CHANGE_LOCALE" : "Change Locale",
            "LOGIN_SESSION_INVALID":"Login session is invalid. Return to login form.",
            "SERVER_ACCESS_PROBLEM":"Unknown problem has occured on server connection. Screen will be reloaded. If this problem persists, please contact the system administrator.",
            "CONFIRM_OK": "Confirm",
            "EDIT":"Edit",
            "CONFIRM_NG": "Cancel",
            "OK": 'OK',
            "LOGIN_SESSION_VALID_TITLE":"Login session is invalid",
            "SERVER_ACCESS_PROBLEM_TITLE": "Unknown problem has occured on server connection",
            "REGISTER":"Register",
            "REMOVE":"Remove",
            "RECEIVED_ORDOER_NO":"Received Order No",
            "ESTIMATIONS_NO":"Estimation No",
            "SAVE":"Save",
            "CLIENT":"Client",
            "ADDRESS_SET":"Address Set",	
            "VALIDATION_ERROR_TITLE": "Error",
            "VALIDATION_ERROR": "There are invalid values in the form.",
            "AREA_MAS": "Received Order No or Estimation No",
            "POST_SUCCESS_TITLE": "Complete",
            "POST_SUCCESS": "Received order address set setting was created.",
            "CONFIRM_ADD_GROUP": "Confirmation",
            "CONFIRM_ADD_GROUP_MESSAGE": "Are you sure you want to register these?",
            "POST_ERROR_STRING_EXIT":"Received Order is not existed. ",
            "POST_ERROR_STRING":"Invoice is issued. ",
            "CONFIRM_REMOVE_GROUP": "Confirmation",
            "CONFIRM_REMOVE_GROUP_MESSAGE": "Are you sure you want to remove these?",
            "REMOVE_SUCCESS_TITLE": "Complete",
            "REMOVE_SUCCESS": "Received order address set setting was remove.",
        },
    }
});
var brosHeader, brosFooter, brosMenu, importFlg

const tryImport = async function() {
    try {
        brosHeader = await import("/header/vue2/x-brostools-header.js")
        brosFooter = await import("/footer/vue2/x-brostools-footer.js")
        brosMenu = await import("/menu/vue2/x-brostools-menu.js")
        importFlg = true
    } catch(e) {
        console.log(e)
        brosHeader = {}
        brosFooter = {}
        brosMenu = {}
        importFlg = false
    }
}
VeeValidate.Validator.localize({
    jp: {
        messages: {
            required: (field) => `${field}は必須です`,
            min: (field, length) => `${field} は ${length} 文字の入力が必要です`,
            max: (field, length) => `${field} は ${length} 文字までしか入力できません`,
            numeric: (field) => `${field} は数字しか入力できません`,
            regex: (field) => `${field} の形式が間違っています`,
            decimal: (field, length) => `${field} は小数第${length}位までしか入力できません`,
            date_format: (field) => `${field} YYYYMMDDの形式である必要があります`
        }
    },
    en: {
        messages: {
            required: (field) => `${field} is a required field`,
            min: (field, length) => `${field} requires a minimum of ${length} characters`,
            max: (field, length) => `${field} allows a maximum of ${length} characters`,
            numeric: (field) => `Only numeric characters are allowed for ${field}`,
            regex: (field) => `${field}'s format is incorrect`,
            decimal: (field, length) => `${field} can only have up to ${length} decimals places`,
            date_format: (field) => `${field} must be in the format YYYYMMDD`
        }
    }
});
Vue.use(VeeValidate);
await tryImport()
var app = new Vue({
    el: "#app_list",
    i18n: i18n,
    data: {
        locale: DEFAULT_LOCALE,
        app_rcvmapping_show: false,
        bros_token: {},
        bros_session: {},
        bros_token_inside:{},
        bros_token_claims: {},
        validationMessages: {
            jp: {
                messages: {
                    required: field => `${field} は必須です`,
                }
            },
            en: {
                messages: {
                    required: field => `${field} is a required field`,
                }
            }
        },
        load_flags: {
            header: false,
            footer: false,
            filter: false,
            menu: false
        },
        loader: $(".vld-overlay"),
        load_err: false,
        rcvordernos:"",
        checkedtype:'register',
        inputtype:'rcv',
        client_cd:"",
        address_set:"",
        address_list:[],
        addrset_id:"",
    },
    components: {
        "x-brostools-header": brosHeader.default,
        "x-brostools-footer": brosFooter.default,
        "x-brostools-menu": brosMenu.default,
    },
    watch: {
    },
    computed: {
        checkLoadStatus() {
            // check component load flags to see if all components are loaded and ready or not
            var loadArr = Object.values(this.load_flags)
            for (var i=0; i<loadArr.length; i++) {
                if (loadArr[i] == "waiting") {
                    return "still loading"
                } else if (loadArr[i] == "error") {
                    if (!this.load_err) {
                        this.load_err = true
                        this.$dialog.alert({
                            title: this.$t("SERVER_ACCESS_PROBLEM_TITLE"),
                            body: this.$t("SERVER_ACCESS_PROBLEM")
                        }, {
                            okText: this.$t("OK")
                        }).then(() => {
                            location.reload()
                        })
                    }
                }
            }
            return "loaded"
        }
    },
    async created() {
        // check component import status
        if (!importFlg) {
            this.load_err = true
            this.$dialog.alert({
                title: this.$t("SERVER_ACCESS_PROBLEM_TITLE"),
                body: this.$t("SERVER_ACCESS_PROBLEM")
            }, {
                okText: this.$t("OK")
            }).then(() => {
                location.reload()
            })
        }
        // prepare here to pass token to header before header's beforeMount
        this.bros_token = this.$cookies.get("bros_token")
        this.setTitle()
        this.onload()
    },
    mounted:function(){
    },
    methods: {
        onload:function(){
            if(this.bros_token == null){
                this.bros_token = ""
            }
            axios.get(URL_API_JWT_VERIFY, {
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${this.bros_token}`,
                }
            })
            .then((response) => {
                //Get Session
                this.bros_token_inside = response.data
                this.getLocale(this.bros_token)
                this.waitForComponentLoad()
                
            })
            .catch(err => {
                console.log(err)
                this.loader.hide()
                this.errorHandling(err.response.status,err.response.data)
            })
        },
        getAddrSetList() {
            this.loader.show();
            this.address_list = [];
            this.addrset_id = null;
            axios.get(`${URL_API_MST_ADDRSET}/${this.client_cd}`, {
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${this.bros_token}`
                },
                params: {
                    latest_only: "true"
                }
            }).then(response => {
                this.address_list = response.data.address_sets
                this.loader.hide();
            }).catch(err => {
                this.loader.hide();
            })
        },
        postAddressSet(){
            this.$validator.validateAll()
            .then(result => {
                if (result === false){
                    this.$dialog.alert({
                        title: this.$t("VALIDATION_ERROR_TITLE"),
                        body: this.$t("VALIDATION_ERROR")
                    }, {
                        okText: this.$t("OK")
                    }).then(() => {
                    })
                    return false;
                }
                else{
                    let newstrings = this.rcvordernos.trim().split(' ').join('').split('\t').join('').replace(/\n+/g, '\n').split("\n");
                    if(this.checkedtype==='register'){
                        var data = {
                            "target_no_list":newstrings,
                            "client_cd":this.client_cd,
                            "addrset_id":this.addrset_id.addrset_id,
                            "addrset_version":this.addrset_id.addrset_version
                        }
                        if(this.inputtype==='rcv'){
                            this.createAddressSet(this.bros_token,data,this.inputtype)
                        }
                        else if(this.inputtype==='est'){
                            this.createAddressSet(this.bros_token,data,this.inputtype)
                        }
                    }
                    else{
                        var data = {
                            "target_no_list":newstrings,
                        }
                        if(this.inputtype==='rcv'){
                            this.removeAddressSet(this.bros_token,data,this.inputtype)
                        }
                        else if(this.inputtype==='est'){
                            this.removeAddressSet(this.bros_token,data,this.inputtype)
                        }
                    }
                }
            });

        },
        createAddressSet(token,data,type) {
            this.$dialog.confirm({
                title: this.$t("CONFIRM_ADD_GROUP"),
                body: this.$t("CONFIRM_ADD_GROUP_MESSAGE"),
            },{
                okText: this.$t("OK"),
                cancelText: this.$t("CONFIRM_NG"),
            })
            .then(() => {
                this.loader.show();
                axios.post(`${URL_API_RCVMAPPING}?by=${type}`, data, {
                    headers: {
                        "Content-Type": "application/json",
                        "Authorization": `Bearer ${token}`,
                    }
                })
                .then(() => {
                    this.loader.hide();
                    this.$dialog.alert({
                        title: this.$t("POST_SUCCESS_TITLE"),
                        body: this.$t("POST_SUCCESS")
                    }, {
                        okText: this.$t("OK")
                    }).then(() => {
                        
                    })
                })
                .catch(err => {
                    this.loader.hide();
                    this.errorHandling (err.response.status,err.response.data)
                })
            })
            .catch((err) => {
            });
        },
        removeAddressSet(token,data,type) {
            this.$dialog.confirm({
                title: this.$t("CONFIRM_REMOVE_GROUP"),
                body: this.$t("CONFIRM_REMOVE_GROUP_MESSAGE"),
            },{
                okText: this.$t("OK"),
                cancelText: this.$t("CONFIRM_NG"),
            })
            .then(() => {
                this.loader.show();
                axios.post(`${URL_API_RCVMAPPING_REMOVE}?by=${type}`, data, {
                    headers: {
                        "Content-Type": "application/json",
                        "Authorization": `Bearer ${token}`,
                    }
                })
                .then(() => {
                    this.loader.hide();
                    this.$dialog.alert({
                        title: this.$t("REMOVE_SUCCESS_TITLE"),
                        body: this.$t("REMOVE_SUCCESS")
                    }, {
                        okText: this.$t("OK")
                    }).then(() => {
                        
                    })
                })
                .catch(err => {
                    this.loader.hide();
                    this.errorHandling (err.response.status,err.response.data)
                })
            })
            .catch((err) => {
            });
        },
        waitForComponentLoad() {
            if (this.checkLoadStatus == "loaded" & !this.load_err) {
                this.emitChangeLocale(this.$i18n.locale)
                this.headerMounted()
                this.loader.hide()
            }
        },
        headerMounted() {
            this.app_rcvmapping_show = true
        },
        emitComponentLoad(result) {
            this.load_flags[result.component] = result.status
        },
        emitChangeLocale(locale) {
            this.$i18n.locale = locale
            if (this.load_flags.footer=="ok") {
                this.$refs.bros_footer.changeLocale(locale)
            }
            if (this.load_flags.menu=="ok") {
                this.$refs.bros_menu.changeLocale(locale)
            }
            this.$validator.localize(locale)
            this.$validator.reset()
            this.setTitle()
        },
        getLocale(token) {
            axios.get(URL_API_USER_LOCALE, {
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${token}`
                }
            }).catch(err => {
                console.log("getLocale error", err.response)
                this.errorHandling(err.response.status,err.response.data)
            })
        },
        setTitle() {
            document.title = "BROS_TOOLS - " + this.$t("PAGE_TITLE")
        },
        getSession(token) {
            axios.get(URL_API_JWT_SESSION, {
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${token}`
                }
            }).then(response => {
                console.log("getSession success: ", response)
                this.bros_session = response.data
            }).catch(err => {
                console.log("getSession error", err.response)
                this.errorHandling(err.response.status,err.response.data)
            })
        },
        setSession(token, session) {
            console.log("session: ", session)
            axios.put(URL_API_JWT_SESSION, session, {
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${token}`
                }
            }).catch(err => {
                console.log("setSession error", err.response)
                this.errorHandling(err.response.status,err.response.data)
            })
        },
        errorHandling(status,data) {
            switch(status) {
                case 401:
                    this.$dialog.alert({
                        title: this.$t("SERVER_ACCESS_PROBLEM_TITLE"),
                        body: this.$t("SERVER_ACCESS_PROBLEM")
                    }, {
                        okText: this.$t("OK")
                    }).then(() => {
                        window.location.href = "/"
                    })
                    break
                case 400:
                    var error_string = ''
                    for(var x in data.errors){
                        let temp = data.errors[x].message.split(": ")
                        let rcv_item = ''
                        if(temp.length>1){
                            rcv_item = temp[1];
                        }
                        if(data.errors[x].code==800){
                            if(this.inputtype==='rcv'){
                                error_string += this.$t("POST_ERROR_STRING_EXIT") + this.$t("RECEIVED_ORDOER_NO") + ":"+ rcv_item+ "<br>";
                            }
                            else{
                                error_string += this.$t("POST_ERROR_STRING_EXIT") + this.$t("ESTIMATIONS_NO") + ":"+ rcv_item+ "<br>";
                            }
                            
                        }
                        else if(data.errors[x].code==801){
                            var mySubString = data.errors[x].message.substring(
                                data.errors[x].message.indexOf(": ") + 1, 
                                data.errors[x].message.lastIndexOf(",")
                            );
                            let temp_rcv = data.errors[x].message.split("rcv_no: ")
                            let rcv_item = temp_rcv[1]
                            
                            if(this.inputtype==='rcv'){
                                error_string += this.$t("POST_ERROR_STRING") + this.$t("RECEIVED_ORDOER_NO") + ": "+ rcv_item+ "<br>";
                            }
                            else{
                                error_string += this.$t("POST_ERROR_STRING") + this.$t("ESTIMATIONS_NO") + ":"+ mySubString+"、"+this.$t("RECEIVED_ORDOER_NO")+":"+rcv_item+ "<br>";
                            }
                        }
                    }
                    this.$dialog.alert({
                        title: this.$t("VALIDATION_ERROR_TITLE"),
                        body: error_string
                    }, {
                        okText: this.$t("OK")
                    }).then(() => {
                        //window.location.href = "/"
                    })
                    break
                default:
                    this.$dialog.alert({
                        title: this.$t("SERVER_ACCESS_PROBLEM_TITLE"),
                        body: this.$t("SERVER_ACCESS_PROBLEM")
                    }, {
                        okText: this.$t("OK")
                    }).then(() => {
                            window.location.href = "/op/rcvmapping"
                    })
                    break
            }
        },

    }
})