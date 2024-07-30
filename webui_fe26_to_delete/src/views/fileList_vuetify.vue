<script setup></script>

<template>
  <v-list density="compact">
    <!--    <v-list-subheader>REPORTS</v-list-subheader>-->

    <v-list-item
      v-for="(item, i) in items"
      :key="i"
      :value="item"
      color="primary"
    >
      <template v-slot:prepend>
        <v-icon icon="fa-solid fa-folder-open" />
      </template>

      <v-list-item-title v-text="item.text"></v-list-item-title>
    </v-list-item>
  </v-list>
</template>
<script>
export default {
  data: () => ({
    items: [
      { text: "Real-Time", icon: "mdi-clock" },
      { text: "Audience", icon: "mdi-account" },
      { text: "Conversions", icon: "mdi-flag" },
    ],
  }),
};
</script>
