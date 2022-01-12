<template>
<div class="container-fluid">
  <div class="navbar-header">
    <div class="logo">
      <a href="/" class="navbar-brand1"><img src="@/assets/images/logo-dark.png"></a>
    </div>
    <div class="menu">
      <ul>
        <li><a href="#">Projects</a></li>
        <li><a href="#">Allocations</a></li>
        <li><a href="#">Dashboard</a></li>
        <li><a href="#">Quests</a></li>
      </ul>
    </div>
    <div class="connect_wallet">
      <p v-if="address == null"><a href="#" class="navbar-brand1" @click="login">Connect Wallet</a></p>
      <p v-if="address !== null"><a href="#" class="navbar-brand1">{{address}}</a></p>
    </div>
  </div>
</div>
</template>
<script>
  import Web3 from 'web3/dist/web3.min.js'
  import AppClient from '@/plugins/client.js';
  const appClient = new AppClient();
  export default {
    name: 'AdminHeader',
    data () {
      return {
        isShow:false,
        isShowLang:false,
        address:null
      }
    },
    created: async function() {
      try {
        await initWeb3();
        this.isLoginButtonDisabled = false;
        this.observe();
        this.address = await getAddress();
      }
      catch (e) {
        console.log(e)
      }
    },
    methods: {
      login: async function() {
        try {
          if (window.ethereum) {
            await ethereum.enable();
          }
          console.log(ethereum)
          let address = await getAddress();
          console.log(address)
          if (address === null) {
            //throw new AppError('Please unlock MetaMask account');
          }

          let networkId = await getNetworkId();
          if (networkId !== 3) {
            //throw new AppError('Please connect MetaMask to Ropsten Test Network');
          }

          let challengeResult = await appClient.challenge(address);
          let signature       = await signTypedData(address, challengeResult.challenge);
          let authorizeResult = await appClient.authorize(address, signature);

          appClient.setToken(authorizeResult.token);

          this.user                 = await appClient.getUser(address);
          this.isObservationEnabled = true;
        }
        catch (e) {
          console.log(e)
        }
      },
      observe: async function() {
        try {
          setTimeout(this.observe, 1000);

          if (!this.isObservationEnabled) {
            return;
          }

          let address = await getAddress();
          if (address === this.user.address) {
            return
          }
          else {
            this.$alert('Account has changed');
            this.logout();
          }
        }
        catch (e) {
          this.handleError(e);
        }
      },
    }
    
  }
  function initWeb3() {
    return new Promise((resolve, reject) => {
      console.log(window.web3)
      if (window.ethereum) {
        window.web3 = new Web3(ethereum);
      }
      else if (window.web3) {
        window.web3 = new Web3(web3.currentProvider);
      }
      else {
        return reject(new AppError('Please install MetaMask'));
      }
      try {
        window.localStorage;
      }
      catch(e) {
        return reject(new AppError('Please allow 3rd party cookies for web3.js 1.0.0'));
      }
      window.web3.extend({
        property: 'app',
        methods: [{
          name: 'signTypedData',
          call: 'eth_signTypedData',
          params: 1,
        }],
      });
      resolve();
  });
}
function getAddress() {
  return new Promise((resolve, reject) => {
    web3.eth.getAccounts().then((accounts) => {
      if (accounts.length <= 0) {
        return resolve(null);
      }
      resolve(accounts[0]);
    })
  });
}
function getNetworkId() {
  return web3.eth.net.getId();
}
function signTypedData(address, value) {
  return web3.app.signTypedData([{
    type: 'string',
    name: 'challenge',
    value: value,
  }], address);
}
</script>