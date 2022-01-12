<template>
    <div id="type-game">
        <canvas id="c-game-canvas" width="900" height="500" ref="canvasElement"></canvas>
    </div>
</template>
<script>
    let score = 0;
let gameFrame = 0;
export default {
    data(){
        return {
            ctx: null,
            canvas:null,
            frameX:0,
            frameY:0,
            spriteWidth :160,
            spriteHeight : 105,
            playerLeft:null
        }
    },
    methods: {
        initGame(){
            this.canvas = document.getElementById('c-game-canvas');
            this.ctx = this.canvas.getContext('2d');
            
            this.playerLeft = new Image();
            this.playerLeft.src = 'https://giayfutsal.com.vn/uploads/pot.png';
            this.ctx.drawImage(this.playerLeft,290,220);

            this.ctx.fillStyle = 'blue';
            this.ctx.beginPath();
            this.ctx.arc(210,260+60, 50, 0, Math.PI * 40);
            this.ctx.fill();
            this.ctx.stroke();
            this.playerLeft = new Image();
            this.playerLeft.src = 'https://giayfutsal.com.vn/uploads/pot.png';
            this.ctx.drawImage(this.playerLeft,210,260);
            //console.log(this.playerLeft)
            //this.playerLeft = new Image();
            //this.playerLeft.src = 'https://giayfutsal.com.vn/uploads/pot.png';
            //this.ctx.drawImage(this.playerLeft,130,300);
            requestAnimationFrame(this.initGame);
        },
    },
    created: function() {
        //this.playerLeft = new Image();
    },
    computed: {
    },
    mounted () {
        this.initGame();
    },
}
</script>

