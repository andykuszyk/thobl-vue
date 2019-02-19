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
        },
        onMouseMove: function(event) {
            for(let obl of this.$children) {
                if(!obl.isDragging) continue;
                obl.onMouseMove(event);
            }
        },
        onMouseDown: function(event) {
            for(let obl of this.$children) {
                if(obl.isOver(event.pageX, event.pageY)) {
                    obl.isDragging = true;
                    obl.select();
                } else {
                    obl.deselect();
                }
            }
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
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  width: 100vw;
  height: 100vh;
  position: absolute;
  left: 0;
  top: 0;
  z-index: -1;
}
</style>
