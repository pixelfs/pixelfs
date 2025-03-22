import {DefaultTheme, defineConfig} from 'vitepress'

export const en= defineConfig({
    lang: 'en',
    description: 'A cross-device file system',

    themeConfig: {
        nav: nav(),
        sidebar: sidebar(),
        editLink: {
            pattern: 'https://github.com/pixelfs/pixelfs/edit/master/docs/:path',
            text: 'Edit this page on GitHub'
        },
        footer: {
            message: 'Released under the GPL-3.0 License.',
            copyright: 'Copyright © 2025-present PixelFS'
        }
    }
})

function nav(): DefaultTheme.NavItem[] {
    return [
        {text: 'Home', link: '/'},
        {text: 'FAQ', link: '/faq'},
    ]
}

function sidebar(): DefaultTheme.Sidebar {
    return [
        {
            text: 'Guide',
            items: [
                {text: 'What is PixelFS?', link: '/'},
                {text: 'Quick Start', link: '/quick-start'},
                {text: 'Configuration', link: '/configuration'},
                {text: 'File Sync', link: '/sync'},
                {text: 'WebDAV', link: '/webdav'},
                {text: 'FAQ', link: '/faq'},
            ]
        },
        {
            text: 'Commands',
            items: [
                {text: 'Auth', link: '/commands/auth'},
                {text: 'Cd', link: '/commands/cd'},
                {text: 'Cp', link: '/commands/cp'},
                {text: 'Daemon', link: '/commands/daemon'},
                {text: 'Download', link: '/commands/download'},
                {text: 'Id', link: '/commands/id'},
                {text: 'Location', link: '/commands/location'},
                {text: 'Ls', link: '/commands/ls'},
                {text: 'M3U8', link: '/commands/m3u8'},
                {text: 'Mkdir', link: '/commands/mkdir'},
                {text: 'Mv', link: '/commands/mv'},
                {text: 'Node', link: '/commands/node'},
                {text: 'Pwd', link: '/commands/pwd'},
                {text: 'Rm', link: '/commands/rm'},
                {text: 'Shell', link: '/commands/shell'},
                {text: 'Storage', link: '/commands/storage'},
                {text: 'StorageLink', link: '/commands/storage-link'},
                {text: 'Sync', link: '/commands/sync'},
                {text: 'Touch', link: '/commands/touch'},
                {text: 'Upload', link: '/commands/upload'},
                {text: 'Version', link: '/commands/version'},
                {text: 'Webdav', link: '/commands/webdav'},
            ]
        },
    ]
}
