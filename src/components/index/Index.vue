<template>
    <div id="type-game">
        <canvas id="c-game-canvas" width="900" height="500" ref="canvasElement" @click="click"></canvas>
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
            playerLeft:null,
            mouse :{
                click:false,
                x:0,
                y:0
            },
            points:[]
        }
    },
    methods: {
        
        initGame(){

            this.canvas = document.getElementById('c-game-canvas');
            this.ctx = this.canvas.getContext('2d');
                       
            
            this.playerLeft = new Image();
            this.playerLeft.src = 'https://giayfutsal.com.vn/uploads/pot.png';
            this.playerLeft.onload = () => this.ctx.drawImage(this.playerLeft,290,220);
            

            this.ctx.fillStyle = 'blue';
            this.ctx.beginPath();
            this.ctx.arc(210,260+60, 70, 0, Math.PI * 40);
            this.ctx.fill();
            //this.ctx.stroke();
            this.playerLeft = new Image();
            this.playerLeft.src = 'https://giayfutsal.com.vn/uploads/pot.png';
            this.playerLeft.onload = () => this.ctx.drawImage(this.playerLeft,210,260);
            
            //console.log(this.playerLeft)
            //this.playerLeft = new Image();
            //this.playerLeft.src = 'https://giayfutsal.com.vn/uploads/pot.png';
            //this.ctx.drawImage(this.playerLeft,130,300);
            //requestAnimationFrame(this.initGame);
        },
        draw: function () {
            var canvas = this.$data.canvas,
            ctx = this.$data.ctx;
            var px = Math.floor(canvas.width / 10);

            // draw circles
            this.points.forEach(function (pt) {
                ctx.beginPath();
                ctx.arc(pt.x, pt.y, 20, 0, Math.PI * 2);
                ctx.fill();
            });
        },
        click: function (e) {
            var bx = e.target.getBoundingClientRect();
            this.points.push({
                x: (e.touches ? e.touches[0].x : e.clientX) - bx.left,
                y: (e.touches ? e.touches[0].y : e.clientY) - bx.top
            });
            if (this.points.length > 10) {
                this.points = this.$data.points.slice(1, 11);
            }
            this.draw();
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

