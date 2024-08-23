<script setup>
import Vertical from '@/components/generic/Vertical.vue'
import DirectoryList from '@/components/directoryList.vue'
import FileList from '@/components/fileList.vue'
import Sidebar from '@/components/generic/simpleSidebar.vue'
import Logo from '@/components/Logo.vue'
import { useFileStore } from '@/stores/files.js'
import { computed, onMounted, ref, watch } from 'vue'
import Error from '@/components/generic/error.vue'
import Breadcrumb from 'primevue/breadcrumb'
import { DateTime } from 'luxon'
import path from 'path-browserify'
import { useRoute, useRouter } from 'vue-router'
import Upload from '@/components/upload.vue'
import { userInfoStore } from '@/stores/info.js'
import Button from 'primevue/button'

const route = useRoute()
const router = useRouter()
const store = useFileStore()
const info = userInfoStore()

// on mount load the content based on the navigation path
onMounted(() => {
    // console.log("on mount, load: "+store.getLocation(route.params.path))
    store.load(store.getLocation(route.params.path))
    info.load()
})
// when a dir is clicked, we change the url
const click = (dirname) => {
    // this joins the vue sub-path "files" with the file system location
    // console.log("on click, push path: "+path.normalize(path.join("/files",store.getLocation(dirname))))
    router.push(path.normalize(path.join('/files', store.getLocation(dirname))))
}

// this watcher watches for changes on the url and appies them
watch(
    () => route.params.path,
    (newLocation, oldId) => {
        store.load(path.normalize(path.join('/', newLocation)))
    }
)

function timeToUnix(input) {
    return DateTime.fromISO(input).toSeconds()
}

const getNames = computed(() => {
    const files = store.files
    const nodes = {
        dirs: [],
        files: []
    }
    for (const [key, value] of Object.entries(files)) {
        if (value.IsDir) {
            nodes.dirs.push({
                key: key + 1,
                label: value.Name,
                selectable: true,
                icon: 'pi pi-fw pi-folder'
            })
        } else {
            nodes.files.push({
                icon: 'pi pi-fw pi-folder',
                name: value.Name,
                filePath: path.join(store.filePath, value.Name),
                size: value.Size,
                modTime: timeToUnix(value.ModTime)
            })
        }
    }
    return nodes
})

const home = ref({
    icon: 'pi pi-home',
    label: 'Root',
    command: () => {
        router.push(path.normalize('/files/'))
    }
})
const breadcrumbItems = computed(() => {
    const items = store
        .getLocation('')
        .split('/')
        .filter(function (el) {
            return el !== ''
        })
    const out = []
    let i = 0
    let base = '/'
    while (i < items.length) {
        base = path.join(base, items[i])
        let p = base
        out.push({
            label: items[i],
            command: () => {
                router.push(path.normalize(path.join('/files/', p)))
            }
        })
        i++
    }
    return out
})

const sidebarOpen = ref(true)
const hclick = () => {
    sidebarOpen.value = !sidebarOpen.value
}

const first = ref('banana')
</script>

<template>
    <Error />

    <Sidebar
        v-model:open="sidebarOpen"
        :breakpoints="{ 1199: 'medium', 600: 'small' }"
        :mobileOpen="false"
    >
        <template v-slot:menu>
            <vertical :center-content="false">
                <template v-slot:header>
                    <Logo />
                    <hr class="space" />
                </template>
                <template v-slot:main>
                    <directory-list :dirs="getNames.dirs" :select="click" />
                </template>
                <template v-slot:footer>
                    <hr class="space" />
                    <div class="fe26-version">
                        Running <a href="https://github.com/AndresBott/Fe26">Fe26 </a>version:
                        {{ info.version }}

                    </div>
                </template>
            </vertical>
        </template>
        <template v-slot:default>
            <vertical>
                <template v-slot:header>
                    <Breadcrumb :home="home" :model="breadcrumbItems" />
                </template>
                <template v-slot:main>
                    <file-list :files="getNames.files" />
                </template>
                <template v-slot:footer>
                    <Upload />
                </template>
            </vertical>
        </template>
    </Sidebar>
</template>
<style lang="scss">
// style the sidebar
.ss-wrapper {
    .ss-left {
        background: var(--bg-color-dark);
        .p-tree {
            background: none;
            width: revert !important;
            padding: 0;
            .p-tree-node-content,
            .p-tree-node-icon,
            span{
                color: var(--bg-color-light);
            }
            .p-tree-node-content:hover {
                background: var(--accent-color-dark);
            }
            .p-tree-node-content:hover span {
                color: var(--bg-color-light);
            }
            .p-tree-node-content {
                transition-duration: 0s;
            }
            .p-tree-root-children {
                gap: 0;
            }
        }
    }
    .ss-main {
        background: var(--bg-color-light);
    }
}

@media (min-width: 900px) {
    :root {
        --ss-width: 35%;
    }
}
@media (min-width: 1200px) {
    :root {
        --ss-width: 420px;
    }
}

.fe26-version {
    font-size: 90%;
    padding: 0.5rem 2rem 1rem 2rem;
    color: var(--accent-color);
}
.fe26-version a {
    font-weight: bold;
    text-decoration: none;
    color: var(--bg-color-light);
}
</style>
