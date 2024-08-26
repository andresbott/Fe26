<template>
    <div class="c-sidebar">
        <div class="c-left">
            <slot name="left"></slot>
        </div>

        <div class="c-main">
            <slot></slot>
        </div>
    </div>
</template>

<style>
.c-sidebar {
    display: grid;
    //grid-template-columns: 450px auto;
    grid-template-rows: 78px auto;
    width: 100%;
    min-height: 100vh;
}
.c-sidebar > .c-left {
    background: var(--bg-color-dark);
}
.c-sidebar > .c-main {
    background: var(--bg-color-light);
}
.c-sidebar > .c-left .p-tree {
    background: none;
    width: revert !important;

    padding: 0;
}

.c-sidebar > .c-left .p-tree .p-tree-node-content,
.c-sidebar > .c-left .p-tree .p-tree-node-icon {
    color: var(--bg-color-light);
}

.c-sidebar > .c-left .p-tree span {
    color: var(--text-color-ligth);
}
.c-sidebar > .c-left .p-tree .p-tree-node-content:hover {
    background: var(--accent-color-dark);
}

.c-sidebar > .c-left .p-tree .p-tree-node-content:hover span {
    color: var(--bg-color-light);
}

.c-sidebar > .c-left .p-tree .p-tree-node-content {
    transition-duration: 0s;
}
.c-sidebar > .c-left .p-tree .p-tree-root-children {
    gap: 0;
}

@media (min-width: 600px) {
    .c-sidebar {
    }
}
@media (min-width: 900px) {
    .c-sidebar {
        grid-template-columns: 35% auto;
        grid-template-rows: revert;
    }
}
@media (min-width: 1200px) {
    .c-sidebar {
        grid-template-columns: 420px auto;
    }
}
</style>
