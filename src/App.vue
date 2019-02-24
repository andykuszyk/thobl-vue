<template>
  <div id="app" v-on:mousemove="onMouseMove" v-on:mouseup="onMouseUp" v-on:mousedown="onMouseDown">
    <Obl v-for="obl in obls" 
        v-bind:radius="obl.radius" 
        v-bind:left="obl.left" 
        v-bind:top="obl.top"
        v-bind:label="obl.label"
        />
  </div>
</template>

<script>
import Obl from './components/Obl.vue'

export default {
    name: 'app',
    data() {
        return {
            isCanvasSelected: false,
            origin: {
                x: 0,
                y: 0
            },
            obls: [ 
                { radius: 100, left: 10, top: 10, label: "small" },
                { radius: 200, left: 100, top: 100, label: "medium" },
                { radius: 300, left: 200, top: 200, label: "ellipse" },
            ]
        }
    },
    methods: {
        onMouseUp: function(event) {
            for(let obl of this.$children) {
                obl.isDragging = false;
            }
            this.isCanvasSelected = false;
        },
        onMouseMove: function(event) {
            if(this.isCanvasSelected) {
                this.origin.x += event.movementX;
                this.origin.y += event.movementY;
                for(let obl of this.$children) {
                    obl.originX = this.origin.x;
                    obl.originY = this.origin.y;
                }
            } else {
                for(let obl of this.$children) {
                    if(!obl.isDragging) continue;
                    obl.onMouseMove(event);
                }
            }
        },
        onMouseDown: function(event) {
            let isOblSelected = false;
            for(let obl of this.$children) {
                if(obl.isOver(event.pageX, event.pageY)) {
                    obl.isDragging = true;
                    isOblSelected = true;
                    obl.select();
                } else {
                    obl.deselect();
                }
            }
            if(isOblSelected) return;
            this.isCanvasSelected = true;
        },
    },
    components: {
        Obl 
    }
}
</script>

<style>
body {
    background: black;
}
#app {
  position: absolute;
  left: 0;
  top: 0;
  right: 0;
  bottom: 0;
  z-index: -1;
  overflow-x: hidden;
  overflow-y: hidden;
}
</style>
