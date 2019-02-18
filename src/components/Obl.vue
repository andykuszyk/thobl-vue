<template>
    <div class="obl-outside"
         v-on:mouseover="onMouseOver"
         v-on:mouseout="onMouseOut"
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
        }
    },
    methods: {
        onMouseOver: function(event) {
            this.highlightColour = 'gray';
        },
        onMouseOut: function(event) {
            this.highlightColour = 'black';
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

