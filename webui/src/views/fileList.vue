<script setup>
  import Tree from 'primevue/tree';
  import { ref, onMounted, computed } from "vue";
  import { useUserStore } from "../stores/users";

  const store = useUserStore();

  const getNames = computed(() =>{
    const users = store.getUsers
    const nodes = [];

    for (const [key, value] of Object.entries(users)) {

      // console.log(value.name)
      nodes.push({
        key: value.id.toString(),
        label: value.name,
        selectable: true,
        icon: "pi pi-fw pi-inbox",
      });
    }
    return nodes;
  })
  const selectedKey = ref(null);
  const onNodeSelect = (node) => {
    console.log(node)
    // toast.add({ severity: 'success', summary: 'Node Selected', detail: node.label, life: 3000 });
  };


  onMounted(() => {
    store.fetchUsers();
  });
</script>

<template>
  <Tree v-model:selectionKeys="selectedKey" :value="getNames"  @nodeSelect="onNodeSelect" selectionMode="single" class="w-full md:w-30rem"></Tree>



</template>


