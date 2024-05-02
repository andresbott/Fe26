<script setup>
import { RouterLink, RouterView } from 'vue-router'
import PageLayout from '@/views/Fe26Layout.vue'
import DirectoryList from "@/components/directoryList.vue";

</script>
<template>
<!--  <RouterView />-->
  <PageLayout>

    <template v-slot:header>
      <nav>
        <!--                <RouterLink to="/">Home</RouterLink>-->
        <!--                <RouterLink to="/about">About</RouterLink>-->
        <!--        Buy VS rent https://www.youtube.com/watch?v=Uwl3-jBNEd4 -->
      </nav>
    </template>
    <template v-slot:menu>
      <directory-list/>
    </template>

  </PageLayout>
</template>
