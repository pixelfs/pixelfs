import { defineConfig } from 'vitepress'
import { en } from './en'
import { zh } from './zh'

// https://vitepress.dev/reference/site-config
export default defineConfig({
    title: "PixelFS",
    description: "A cross-device file system",
    head: [
        ['link', { rel: 'icon', type: 'image/png', href: '/logo.png' }],
        ['meta', { property: 'og:type', content: 'website' }],
        ['meta', { property: 'og:title', content: 'PixelFS | A cross-device file system' }],
        ['meta', { property: 'og:site_name', content: 'PixelFS' }],
    ],
    locales: {
        root: {
            label: 'English',
            ...en
        },
        'zh-hans': {
            label: '简体中文',
            ...zh
        }
    },
    themeConfig: {
        // https://vitepress.dev/reference/default-theme-config
        search: {
            provider: "local",
        },
        logo: "/logo.png",
        socialLinks: [
            { icon: 'github', link: 'https://github.com/pixelfs/pixelfs' },
        ],
    },
})
