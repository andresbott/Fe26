<template>
    <div class="c-horizontal" :class="['c-main',  { 'hide-center': hideCenter }]">
        <div class="c-left">
            <slot name="left"></slot>
        </div>

        <div :class="['c-main', { 'h-center': centerContent }]">
            <slot></slot>
        </div>

        <div class="c-right">
            <slot name="right"></slot>
        </div>
    </div>
</template>

<script setup>
defineProps({
    centerContent: {
        type: Boolean,
        default: false
    },
    // if set to true the center content will be 0px
    hideCenter:{
        type: Boolean,
        default: false
    }
})
</script>

<style lang="scss">
.c-horizontal {
    display: grid;
    grid-template-columns: auto 1fr auto;
    width: 100%;

    &> .c-main.h-center {
        display: grid;
        place-items: normal center;
    }
    &.hide-center{
        grid-template-columns: auto 0 auto;
    }
    .c-right{
        display: grid;
        place-items: normal right;
    }

}

</style>
