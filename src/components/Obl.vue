<template>
    <div class="obl-outside"
         v-on:mouseover="onMouseOver"
         v-on:mouseout="onMouseOut"
         v-on:mouseup="onMouseUp"
         v-bind:style="{ 
            width: radius * 2 + 'px', 
            height: radius * 2 + 'px', 
            left: left + 'px', 
            top: top + 'px',
            background: highlightColour,
        }">
        <div class="obl-inside"
             v-bind:style="{
                width: radius * 2 - borderSize + 'px',
                height: radius * 2 - borderSize + 'px',
                left: borderSize / 2 + 'px',
                top: borderSize / 2 + 'px',
                fontSize: radius * 2 + '%',
            }">
            <div class="obl-contents">
                <span>{{ label }}</span>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            borderSize: 10,
            highlightColour: 'black',
            isDragging: false,
            isActive: false,
            mouseOffsetX: null,
            mouseOffsetY: null,
        }
    },
    methods: {
        onMouseOver: function(event) {
            this.highlightColour = 'gray';
        },
        onMouseOut: function(event) {
            this.highlightColour = 'black';
        },
        onMouseMove: function(event) {
            if(!this.isDragging) return;
            if(this.mouseOffsetX == null && this.mouseOffsetY == null) {
                this.mouseOffsetX = event.pageX - this.left;
                this.mouseOffsetY = event.pageY - this.top;
            }
            this.left = event.pageX - this.mouseOffsetX;
            this.top = event.pageY - this.mouseOffsetY;
        },
        onMouseUp: function(event) {
            this.mouseOffsetX = null;
            this.mouseOffsetY = null;
        },
        isOver: function(x, y) {
            if(x < this.left) return false;
            if(y < this.top) return false;
            if(x > this.left + this.radius * 2) return false;
            if(y > this.top + this.radius * 2) return false;
            return true;
        }
    },
    props: ['radius', 'left', 'top', 'label' ],
}
</script>

<style scoped>
.obl-outside {
    position: absolute;
    border-radius: 50%;
}
.obl-inside {
    position: absolute;
    background: white;
    border-radius: 50%;
}
.obl-contents {
    position: relative;
    top: 50%;
    transform: translateY(-50%);
    text-align: center;
    user-select: none;
}
</style>

