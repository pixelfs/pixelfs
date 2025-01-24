import {DefaultTheme, defineConfig} from 'vitepress'

export const zh = defineConfig({
    lang: 'zh-Hans',
    description: '一个跨设备的文件管理系统',

    themeConfig: {
        nav: nav(),
        sidebar: sidebar(),
        editLink: {
            pattern: 'https://github.com/pixelfs/pixelfs/edit/master/docs/:path',
            text: '在 GitHub 上编辑此页面'
        },
        footer: {
            message: '基于 GPL-3.0 许可发布',
            copyright: '版权所有 © 2025-现在 PixelFS'
        },
        docFooter: {
            prev: '上一页',
            next: '下一页'
        },
        outline: {
            label: '页面导航'
        },
    }
})

function nav(): DefaultTheme.NavItem[] {
    return [
        {text: '首页', link: '/zh-hans/'},
        {text: '常见问题', link: '/zh-hans/faq'},
    ]
}

function sidebar(): DefaultTheme.Sidebar {
    return [
        {
            text: '入门',
            items: [
                {text: '什么是 PixelFS?', link: '/zh-hans/'},
                {text: '快速入门', link: '/zh-hans/quick-start'},
                {text: '配置', link: '/zh-hans/configuration'},
                {text: 'WebDAV 服务', link: '/zh-hans/webdav'},
                {text: '常见问题', link: '/zh-hans/faq'},
            ]
        },
        {
            text: '命令参考',
            items: [
                {text: 'Auth', link: '/zh-hans/commands/auth'},
                {text: 'Cd', link: '/zh-hans/commands/cd'},
                {text: 'Cp', link: '/zh-hans/commands/cp'},
                {text: 'Daemon', link: '/zh-hans/commands/daemon'},
                {text: 'Download', link: '/zh-hans/commands/download'},
                {text: 'Id', link: '/zh-hans/commands/id'},
                {text: 'Location', link: '/zh-hans/commands/location'},
                {text: 'Ls', link: '/zh-hans/commands/ls'},
                {text: 'M3U8', link: '/zh-hans/commands/m3u8'},
                {text: 'Mkdir', link: '/zh-hans/commands/mkdir'},
                {text: 'Mv', link: '/zh-hans/commands/mv'},
                {text: 'Node', link: '/zh-hans/commands/node'},
                {text: 'Pwd', link: '/zh-hans/commands/pwd'},
                {text: 'Rm', link: '/zh-hans/commands/rm'},
                {text: 'Shell', link: '/zh-hans/commands/shell'},
                {text: 'Storage', link: '/zh-hans/commands/storage'},
                {text: 'StorageLink', link: '/zh-hans/commands/storage-link'},
                {text: 'Touch', link: '/zh-hans/commands/touch'},
                {text: 'Upload', link: '/zh-hans/commands/upload'},
                {text: 'Version', link: '/zh-hans/commands/version'},
                {text: 'Webdav', link: '/zh-hans/commands/webdav'},
            ]
        },
    ]
}
