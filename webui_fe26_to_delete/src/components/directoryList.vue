<script setup>
import Tree from "primevue/tree";
import { onMounted, computed } from "vue";
import { useFilesStore } from "../stores/files.js";

const store = useFilesStore();

const getNames = computed(() => {
  const files = store.getFiles;
  const nodes = [];

  if (!store.isRoot) {
    nodes.push({
      key: 0,
      label: "..",
      selectable: true,
      icon: "pi pi-fw pi-chevron-up ",
    });
  }

  for (const [key, value] of Object.entries(files)) {
    if (value.IsDir) {
      nodes.push({
        key: key + 1,
        label: value.Name,
        selectable: true,
        icon: "pi pi-fw pi-folder",
      });
    }
  }

  return nodes;
});

const onNodeSelect = (node) => {
  store.Navigate(node.label);
};

onMounted(() => {
  store.fetchFiles("/");
});
</script>

<template>
  <Tree
    :value="getNames"
    @nodeSelect="onNodeSelect"
    selectionMode="single"
    class="w-full md:w-30rem"
  ></Tree>
</template>
