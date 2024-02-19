<script setup>
import Fe26Layout from "@/views/Fe26Layout.vue";
import fileList from "@/views/fileList.vue";
</script>

<template>
  <fe26-layout >
    <template v-slot:menu>
      <file-list></file-list>
      <!-- content for the header slot -->
    </template>
    <div>
      content
    </div>
  </fe26-layout>
</template>

<script>
</script>

<style>
#app,html,body{
  height:100%;
  width:100%
}
</style>