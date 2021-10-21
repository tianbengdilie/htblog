export const dirs = [
    {
        id: 1,
        name: 'dir1',
        fullpath: 'dir1',
        folder: true,
        children: [
            {
                id: 11,
                name: 'file1',
                fullpath: 'dir1/file1',
                folder: false
            },
            {
                id: 12,
                name: 'file2',
                fullpath: 'dir1/file2',
                folder: false
            },
            {
                id: 13,
                name: 'file3',
                fullpath: 'dir1/file3',
                folder: false
            },
            {
                id: 14,
                name: 'dir11',
                fullpath: 'dir1/dir11',
                folder: true,
                children: [
                    {
                        id: 141,
                        name: "file141",
                        fullpath: 'dir1/dir11/file141',
                        folder: true
                    }
                ]
            }
        ]
    },
    { //
        id: 2,
        name: "file20", //
        fullpath: 'file20',
        folder: false
    }
]