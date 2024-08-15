import { definePreset } from '@primevue/themes'
import Lara from '@primevue/themes/lara'

const CustomTheme = definePreset(Lara, {
    primitive: {
        borderRadius: {
            none: '0',
            xs: '1px',
            sm: '1px',
            md: '1px',
            lg: '1px',
            xl: '1px'
        }
    },
    semantic: {
        primary: {
            50: '{blue.50}',
            100: '{blue.100}',
            200: '{blue.200}',
            300: '{blue.300}',
            400: '{blue.400}',
            500: '{blue.500}',
            600: '{blue.600}',
            700: '{blue.700}',
            800: '{blue.800}',
            900: '{blue.900}',
            950: '{blue.950}'
        }
    }
})

export default CustomTheme