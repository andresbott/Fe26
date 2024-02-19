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
      });
    }
    console.log(nodes)
    return nodes;
  })
  onMounted(() => {
    store.fetchUsers();
  });
</script>

<template>
    <Tree :value=getNames class="w-full">    </Tree>

</template>


<script>

import { ref, onMounted } from 'vue';


// import { NodeService } from '@/service/NodeService';

// const nodes = ref(null);

// onMounted(() => {
//   NodeService.getTreeNodes().then((data) => (nodes.value = data));
// });

const nodes = [
  {
    key: "a",
    label: "test1",
    selectable: true,
  },
  {
    key: "b",
    label: "test2",
  },
]

</script>

