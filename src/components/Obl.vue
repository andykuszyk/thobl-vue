<template>
    <div class="obl-outside"
         v-if="!isDeleted"
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
        <div class="obl-delete"
             v-if="isActive"
             v-on:mousedown="onDelete"
             v-bind:style="{
                width: 35 + 'px',
                height: 35 + 'px',
                left: deleteCoords.x + 'px',
                top: deleteCoords.y + 'px',
            }">
            <b>x</b>
        </div>
        <div class="obl-inside"
             v-bind:style="{
                width: radius * 2 - borderSize + 'px',
                height: radius * 2 - borderSize + 'px',
                left: borderSize / 2 + 'px',
                top: borderSize / 2 + 'px',
            }">
            <div class="obl-contents" v-on:mousedown="onMouseDown">
                <span v-if="!isEditing" v-bind:style="{fontSize: radius * 2 + '%'}">{{ labelText }}</span>
                <input v-if="isEditing" 
                       type="text" 
                       v-bind:style="{
                           width: (radius * 2 - borderSize * 2) * 0.9 + 'px',
                           fontSize: radius * 2 + '%',
                           textAlign: 'center',
                           background: 'transparent',
                           color: 'white'
                       }" 
                       v-on:keydown="onKeydown"
                       v-model="labelText"
                       >
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
            isEditing: false,
            labelText: this.label,
            radius: this.initialRadius,
            isDeleted: false,
            deleteCoords: { x: 0, y: 0 },
        }
    },
    props: ['initialRadius', 'left', 'top', 'label', 'id'],
    methods: {
        onDelete: function(event) {
            this.isDeleted = true;
        },
        select: function() {
            this.isActive = true;
            this.mouseOffsetX = null;
            this.mouseOffsetY = null;
            this.calculateOrbitalCoordinates();
        },
        deselect: function() {
            this.isActive = false;
            this.highlightColour = 'white';
            this.isEditing = false;
            if(this.labelText == '') this.labelText = 'Obl';
        },
        calculateOrbitalCoordinates: function() {
            this.deleteCoords = {
                x: Math.cos(Math.PI / 4) * this.radius + this.radius,
                y: this.radius - Math.sin(Math.PI / 4) * this.radius
            };

        },
        onWheel: function(event) {
            if(event.deltaY < 0) {
                this.radius -= this.scaleAmount;
                this.x += this.scaleAmount;
                this.y += this.scaleAmount;
            } else {
                this.radius += this.scaleAmount;
                this.x -= this.scaleAmount;
                this.y -= this.scaleAmount;

            }
            this.calculateOrbitalCoordinates();
            event.stopPropagation();
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
            this.x = event.pageX - this.mouseOffsetX - this.originX;
            this.y = event.pageY - this.mouseOffsetY - this.originY;
        },
        onMouseUp: function(event) {
            this.mouseOffsetX = null;
            this.mouseOffsetY = null;
        },
        isOver: function(x, y) {
            let center = {
                x: this.getAbsolutePosition().x + this.radius,
                y: this.getAbsolutePosition().y + this.radius
            }
            /*
                          +-#
                         # \|  #
                        #   x   #
                         #     #
                            #
            */
            let hypoteneuse = Math.sqrt((center.x - x) ** 2 + (center.y - y) ** 2);
            return hypoteneuse <= this.radius;
        },
        onMouseDown: function(event) {
            this.isEditing = true;
        },
        onKeydown: function(event) {
            if(event.key == "Enter" || event.key == "Escape") {
                this.isEditing = false;
                if(this.labelText == '') this.labelText = 'Obl';
            }
        },
    }
}
</script>

<style scoped>
.obl-outside {
    position: absolute;
    border-radius: 50%;
    z-index: 0;
}
.obl-delete {
    position: absolute;
    background: red;
    color: white;
    border-radius: 50%;
    text-align: center;
    z-index: 2;
}
.obl-inside {
    position: absolute;
    background: black;
    border-radius: 50%;
    z-index: 1;
}
.obl-contents {
    z-index: 2;
    color: white;
    position: relative;
    top: 50%;
    transform: translateY(-50%);
    text-align: center;
    user-select: none;
}
</style>

