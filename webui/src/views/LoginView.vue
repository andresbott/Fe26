<script setup>
import Vertical from '@/components/generic/Vertical.vue'
import Horizontal from '@/components/generic/horizontal.vue'
import UserLogin from '@/components/user/UserLogin.vue'
</script>

<template>
    <vertical :center-content="true">
        <template v-slot:main>
            <horizontal :center-content="true">
                <template v-slot:default>
                    <UserLogin />
                </template>
            </horizontal>
        </template>

        <template v-slot:footer> </template>
    </vertical>
</template>
