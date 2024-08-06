<script setup>
import Vertical from '@/components/legos/Vertical.vue'
import Horizontal from '@/components/legos/horizontal.vue'
import DirectoryList from '@/components/directoryList.vue'
import FileList from '@/components/fileList.vue'
import { useFileStore } from '@/stores/files.js'
import { computed, onMounted } from 'vue'

const store = useFileStore()
onMounted(() => {
    store.goTo('')
})

const getNames = computed(() => {
    const files = store.files
    const nodes = {
        dirs: [],
        files: []
    }

    if (!store.isRoot) {
        nodes.dirs.push({
            key: 0,
            label: '..',
            selectable: true,
            icon: 'pi pi-fw pi-chevron-up '
        })
    }

    for (const [key, value] of Object.entries(files)) {
        if (value.IsDir) {
            nodes.dirs.push({
                key: key + 1,
                label: value.Name,
                selectable: true,
                icon: 'pi pi-fw pi-folder'
            })
        }else {
            nodes.files.push({
                key: key + 1,
                label: value.Name,
                selectable: true,
                icon: 'pi pi-fw pi-folder'
            })
        }
    }

    return nodes
})

</script>

<template>
    <vertical :center-content="false">
        <template v-slot:header>
            <!--            <TopBar />-->
        </template>

        <template v-slot:main>
            <horizontal>
                <template v-slot:left>
                    <directory-list :dirs="getNames.dirs"/>
                </template>
                <template v-slot:default>
                    <file-list :files="getNames.files"/>
                </template>
            </horizontal>
        </template>
        <template v-slot:footer>my footer </template>
    </vertical>
</template>
