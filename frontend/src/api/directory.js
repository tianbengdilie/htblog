export const dirs = [
    {
        id: 1,
        name: 'dir1',
        folder: true,
        children: [
            {
                id: 11,
                name: 'file1',
                folder: false
            },
            {
                id: 12,
                name: 'file2',
                folder: false
            },
            {
                id: 13,
                name: 'file3',
                folder: false
            },
            {
                id: 14,
                name: 'dir11',
                folder: true,
                children: [
                    {
                        id: 141,
                        name: "file141",
                        folder: true
                    }
                ]
            }
        ]
    },
    { //
        id: 2,
        name: "file20", //
        folder: false
    }
]