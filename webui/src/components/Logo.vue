<script setup>

</script>
<style>
.fe26-logo-container{
    margin: 1rem;
    text-align: center;
}
#fe26-logo{
    display: inline-block;
    position: relative;
    background: var(--main-color);
    width: 3rem;
    height: 3rem;
    border-radius: .2rem;
}
#fe26-logo::before {
    content: "Fe";
    text-align: center;
    font-weight: 700;
    font-size: 1.7rem;
    position: absolute;
    left: .5rem;
    bottom: 0.3rem;
    color: var(--bg-color-dark);
}
#fe26-logo::after {
    content: "26";
    font-size: 0.8rem;
    right: .3rem;
    top: .2rem;
    position: absolute;
    color: var(--bg-color-dark);
}
.fe26-text-fe {
    font-size: 1.5rem;
    color: #e6e6e6;
    margin-right: 1.5rem;
    margin-left: 1rem;
}

</style>
<template>
    <div class="fe26-logo-container">
        <div id="fe26-logo"></div>
        <span class="fe26-text-fe">File Explorer</span></div>
</template>
