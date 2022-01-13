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
            points:[],
            plots:[
                {
                    x:290,
                    y:220,
                },
                {
                    x:190,
                    y:260,
                }
            ]
        }
    },
    methods: {
        
        initGame(){
            this.canvas = document.getElementById('c-game-canvas');
            this.ctx = this.canvas.getContext('2d');
            for(let i = 0; i< this.plots.length;i++)
            {
                //this.ctx.fillStyle = 'blue';
                //this.ctx.beginPath();
                //this.ctx.arc(this.plots[i].x+68,this.plots[i].y+34, 40, 0, Math.PI * 40);
                //this.ctx.fill();
                //this.ctx.stroke();
                this.playerLeft = new Image();
                this.playerLeft.src = 'https://giayfutsal.com.vn/uploads/pot.png';
                this.playerLeft.onload = () => this.ctx.drawImage(this.playerLeft,this.plots[i].x,this.plots[i].y);
            }
            this.ctx.restore();
        },
        draw: function () {
            var canvas = this.$data.canvas,
            ctx = this.$data.ctx;
            var px = Math.floor(canvas.width / 10);

            // draw circles
            this.points.forEach(function (pt) {
                ctx.beginPath();
                ctx.arc(pt.x, pt.y, 35, 0, Math.PI * 2);
                ctx.fill();
            });
        },
        click: function (e) {
            console.log(e)
            for(let i = 0; i< this.plots.length;i++)
            {
                let dx = e.layerX - (this.plots[i].x+68);
                let dy = e.layerY - (this.plots[i].y+24);
                let distance = Math.sqrt(dx * dx + dy * dy);
                if(distance<40){
                    this.playerLeft = new Image();
                    this.playerLeft.src = 'https://giayfutsal.com.vn/uploads/pot1.png';
                    this.playerLeft.onload = () => this.ctx.drawImage(this.playerLeft,this.plots[i].x,this.plots[i].y);
                }
            }
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

