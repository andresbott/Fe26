<script setup>
import Vertical from '@/components/legos/Vertical.vue'
import DirectoryList from '@/components/directoryList.vue'
import FileList from '@/components/fileList.vue'
import Sidebar from '@/views/parts/sidebar.vue'
import Logo from '@/components/Logo.vue'
import { useFileStore } from '@/stores/files.js'
import { computed, onMounted, ref, watch } from 'vue'
import Error from '@/views/parts/error.vue'
import Breadcrumb from 'primevue/breadcrumb'
import { DateTime } from 'luxon'
import path from 'path-browserify'
import { useRoute,useRouter } from 'vue-router'

const route = useRoute();
const router = useRouter()
const store = useFileStore()

// on mount load the content based on the navigation path
onMounted(() => {
    // console.log("on mount, load: "+store.getLocation(route.params.path))
    store.load(store.getLocation(route.params.path))
})
// when a dir is clicked, we change the url
const click = (dirname ) =>{
    // this joins the vue sub-path "files" with the file system location
    // console.log("on click, push path: "+path.normalize(path.join("/files",store.getLocation(dirname))))
    router.push(path.normalize(path.join("/files",store.getLocation(dirname))))
}

// this warcher watches for changes on the url and appies them
watch(
    () => route.params.path,
    (newLocation, oldId) => {
        store.load(path.normalize(path.join("/",newLocation)))
    }
)


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
    icon: 'pi pi-home',
    label: "Root",
    command: () =>{
        router.push(path.normalize("/files/"))
    }
});
const breadcrumbItems = computed(()=>{
    const items = store.getLocation("").split('/').filter(function (el) {
        return el !== "";
    });
    const out = [];
    let i = 0;
    let base ="/"
    while (i < items.length) {
        base = path.join(base, items[i])
        let p = base
        out.push( {
            label: items[i] ,
            command: () =>{
                router.push(path.normalize(path.join("/files/",p)))
            }
        },)
        i++;
    }
    return out
})

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
                    <directory-list :dirs="getNames.dirs" :select="click"/>
                </template>
                <template v-slot:footer>my footer </template>
            </vertical>
        </template>
        <template v-slot:default>
            <Breadcrumb :home="home" :model="breadcrumbItems" />
            <file-list :files="getNames.files"/>
        </template>
    </Sidebar>

</template>
