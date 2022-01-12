<template>
    <div id="login">
        <form>
            <h1>Sign In</h1>
            <input type="text" v-model="email" placeholder="Username">
	        <input type="password"  v-model="password" placeholder="Password">
	        <button type="submit" class="btn btn-default" @click="handleSubmit">Sign in</button>
        </form>
    </div>
</template>
<script>
    export default {
        
        data(){
            return {
                email : "",
                password : ""
            }
        },
        methods : {
            handleSubmit(e){
                e.preventDefault()
                if (this.password.length > 0) {
                    this.$http.post('api/user/login', {
                        username: this.email,
                        password: this.password
                    })
                    .then(response => {
                        console.log(response.data.userinfo)
                        if(response.data.userinfo){
                            console.log("1")
                            localStorage.setItem('user',JSON.stringify(response.data.userinfo));
                            localStorage.setItem('token',response.data.token);
                            if (localStorage.getItem('token') != null){
                                this.$emit('loggedIn')
                                if(this.$route.params.nextUrl != null){
                                    this.$router.push(this.$route.params.nextUrl)
                                }
                                else {
                                    this.$router.push('listproject');
                                }
                            }                 
                        }
                    })
                    .catch(function (error) {
                        console.error(error.response);
                    });
                }
            }
        }
    }
</script>
