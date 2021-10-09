/** When your routing table is too long, you can split it into small modules */

import Layout from '@/views/layout';

const permissionRouter = {
    path: '/permission',
    component: Layout,
    redirect: 'noredirect',
    name: 'Permission',
    meta: {
        title: 'route.permissions',
        icon: 'mdi-account-group',
    },
    children: [
        {
            path: 'admin',
            component: () => import('@/components/permission/Admin.vue'),
            name: 'PermissionAdmin',
            meta: { title: 'route.permission.admin', roles: ['admin'], noCache: true },
        },
        {
            path: 'editor',
            component: () => import('@/components/permission/Editor.vue'),
            name: 'PermissionEditor',
            meta: { title: 'route.permission.editor', roles: ['editor'], noCache: true },
        },
        {
            path: 'visitor',
            component: () => import('@/components/permission/Visitor.vue'),
            name: 'PermissionVisitor',
            meta: { title: 'route.permission.visitor', roles: ['visitor'], noCache: true },
        },
    ],
};

export default permissionRouter;
