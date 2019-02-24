<template>
    <div class="obl-outside"
         v-on:mouseover="onMouseOver"
         v-on:mouseout="onMouseOut"
         v-on:mouseup="onMouseUp"
         v-on:wheel="onWheel"
         v-bind:style="{ 
            width: radius * 2 + 'px', 
            height: radius * 2 + 'px', 
            left: originX + x + 'px', 
            top: originY + y + 'px',
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
            x: this.left,
            y: this.top,
            originX: 0,
            originY: 0,
            borderSize: 10,
            highlightColour: 'white',
            isDragging: false,
            isActive: false,
            mouseOffsetX: null,
            mouseOffsetY: null,
            scaleAmount: 5,
        }
    },
    props: ['radius', 'left', 'top', 'label' ],
    methods: {
        select: function() {
            this.isActive = true;
            this.mouseOffsetX = null;
            this.mouseOffsetY = null;
        },
        deselect: function() {
            this.isActive = false;
            this.highlightColour = 'white';
        },
        onWheel: function(event) {
            if(event.deltaY < 0) {
                this.radius -= this.scaleAmount;
                this.x += this.scaleAmount / 2;
                this.y += this.scaleAmount / 2;
            } else {
                this.radius += this.scaleAmount;
                this.x -= this.scaleAmount / 2;
                this.y -= this.scaleAmount / 2;

            }
            event.stopPropogation();
        },
        onMouseOver: function(event) {
            this.highlightColour = 'blue';
        },
        onMouseOut: function(event) {
            if(!this.isActive) this.highlightColour = 'white';
        },
        getAbsolutePosition() {
            return {
                x: this.x + this.originX,
                y: this.y + this.originY
            }
        },
        onMouseMove: function(event) {
            if(!this.isDragging) return;
            if(this.mouseOffsetX == null && this.mouseOffsetY == null) {
                this.mouseOffsetX = event.pageX - this.getAbsolutePosition().x;
                this.mouseOffsetY = event.pageY - this.getAbsolutePosition().y;
            }
            console.log('(' + this.mouseOffsetX + ', ' + this.mouseOffsetY + ')');
            this.x = event.pageX - this.mouseOffsetX - this.originX;
            this.y = event.pageY - this.mouseOffsetY - this.originY;
        },
        onMouseUp: function(event) {
            this.mouseOffsetX = null;
            this.mouseOffsetY = null;
        },
        isOver: function(x, y) {
            if(x < this.getAbsolutePosition().x) return false;
            if(y < this.getAbsolutePosition().y) return false;
            if(x > this.getAbsolutePosition().x + this.radius * 2) return false;
            if(y > this.getAbsolutePosition().y + this.radius * 2) return false;
            return true;
        }
    },
}
</script>

<style scoped>
.obl-outside {
    position: absolute;
    border-radius: 50%;
}
.obl-inside {
    position: absolute;
    background: black;
    border-radius: 50%;
}
.obl-contents {
    color: white;
    position: relative;
    top: 50%;
    transform: translateY(-50%);
    text-align: center;
    user-select: none;
}
</style>

