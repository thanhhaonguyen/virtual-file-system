import React, {useEffect, useState} from "react";
import Terminal from "terminal-in-react";
import CallApi from "../../Api/callApi";
import pseudoFileSystemPlugin from "terminal-in-react-pseudo-file-system-plugin";

const FileSystem = () => {

    // const [folders, setFolders] = useState([])
    // const [currentFolderParentId, setCurrentFolderParentId] = useState(0)
    //
    // useEffect(() => {
    //     CallApi(`folder-by-parent/${currentFolderParentId}`, "GET", "").then(res => {
    //         if (res.data.status === 200) {
    //             setFolders(res.data.data);
    //         }
    //     }).catch(err => {
    //         console.log(err)
    //     })
    // }, [])
    //
    // const ChangeCurrentWorkingDir = (args, print) => {
    //     const dirName = args.slice(1).join(" ")
    //     let folderId = ""
    //     console.log(folders)
    //     folders.map(f => {
    //         if (f.name === dirName) {
    //             folderId = f.id
    //         }
    //     })
    //     CallApi(`folder-by-parent/${folderId}`, "GET", "").then(res => {
    //         if (res.data.status === 200) {
    //             res.data.data.map(e => {
    //                 print(e.name)
    //             })
    //             setFolders(res.data.data)
    //             setCurrentFolderParentId(folderId)
    //         }
    //     }).catch(err => {
    //         print(err.response.data.errors)
    //     })
    // }

    const GetDir = (args, print) => {
        CallApi("folder", "GET", "").then(res => {
            if (res.data.status === 200) {
                res.data.data.map(e => {
                    print(e.name)
                })
            }
        }).catch(err => {
            print(err.response.data.errors)
        })
    }

    const CreateDir = (args, print) => {
        const dirName = args.slice(1)[0]
        const data = {
            name: dirName
        }
        CallApi("folder", "POST", data).then(res => {
            if (res.data.status === 200) {
                console.log("success")
            }
        }).catch(err => {
            print(err.response.data.errors)
        })
    }

    const CreateFile = (args, print) => {
        const fileName = args.slice(1)[0]
        const fileData = args.slice(2)[0]
        const data = {
            name: fileName,
            data: fileData
        }
        CallApi("file", "POST", data).then(res => {
            if (res.data.status === 200) {
                console.log("success")
            }
        }).catch(err => {
            print(err.response.data.errors)
        })
    }

    return (
        <div
            style={{
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                height: "100vh"
            }}
        >
                <Terminal
                    color="green"
                    backgroundColor="black"
                    barColor="black"
                    style={{ fontWeight: "bold", fontSize: "1em" }}
                    commands={{
                        // "cd": (args, print) => ChangeCurrentWorkingDir(args, print),
                        "ls": (args, print) => GetDir(args, print),
                        "mkdir": (args, print) => CreateDir(args, print),
                        "cr": (args, print) => CreateFile(args, print)
                    }}
                    descriptions={{
                        // "cd": "change current working directory",
                        "ls": "list out all items directly under a folder",
                        "mkdir": "create a new directory",
                        "cr": "create a new file"
                    }}
                    msg="Welcome to Virtual-file-system !!!"
                />

        </div>
    )
}

export default FileSystem