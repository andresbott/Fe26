<script setup>
import Vertical from '@/components/legos/Vertical.vue'
import DirectoryList from '@/components/directoryList.vue'
import FileList from '@/components/fileList.vue'
import Sidebar from '@/views/parts/sidebar.vue'
import Logo from '@/components/Logo.vue'
import { useFileStore } from '@/stores/files.js'
import { computed, onMounted, ref } from 'vue'
import Error from '@/views/parts/error.vue'
import Breadcrumb from 'primevue/breadcrumb'
import { DateTime } from 'luxon'
import path from 'path-browserify'



const store = useFileStore()
onMounted(() => {
    store.goTo('')
})

function timeToUnix(input ){
    return DateTime
        .fromISO(input)
        .toSeconds()
}

const getNames = computed(() => {
    const files = store.files
    const nodes = {
        dirs: [],
        files: []
    }

    if (!store.isRoot()) {
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
                icon:  'pi pi-fw pi-folder',
                name: value.Name,
                filePath : path.join(store.filePath,value.Name),
                size: value.Size,
                modTime: timeToUnix(value.ModTime),
            })
        }
    }
    return nodes
})


const home = ref({
    icon: 'pi pi-home'
});
const items = ref([
    { label: 'Electronics' },
    { label: 'Computer' },
    { label: 'Accessories' },
    { label: 'Keyboard' },
    { label: 'Wireless' }
]);
</script>


<template>

    <Error/>
    <Sidebar>
        <template v-slot:left>
            <vertical :center-content="false">
                <template v-slot:header>
                    <Logo/>
                    <hr class="space"/>
                </template>
                <template v-slot:main>
                    <directory-list :dirs="getNames.dirs"/>
                </template>
                <template v-slot:footer>my footer </template>
            </vertical>
        </template>
        <template v-slot:default>
            <Breadcrumb :home="home" :model="items" />
            <file-list :files="getNames.files"/>
        </template>
    </Sidebar>

</template>
