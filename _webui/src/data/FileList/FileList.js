export class FileList {

    constructor(data) {
        this.count = data.Count;
        this.files = [];
        this.dirs = [];
        for (let i = 0; i < data.Items.length; i++) {
            if (data.Items[i].IsDir) {
                this.dirs.push(data.Items[i].Name)
            } else {
                this.files.push(data.Items[i].Name)
            }
        }
    }

    get Files() {
        return this.files
    }

    get Dirs() {
        return this.dirs
    }
}
