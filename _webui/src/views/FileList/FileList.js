import './FileList.css';
import Dirs from "../../components/Dirs/Dirs";
import {usePath} from "../../data/FileList/FileListStore";
import {useEffect} from "react";
import Notification from "../../components/Notification/Notification";


function App() {
    // initial state
    const loadPath = usePath(state => state.loadPath)
    useEffect(() => {
        loadPath("/")
    }, []);


    return (
        <div className="FileExplorer">
            {/*<Notification/>*/}
            <header>
                head
            </header>
            <div className="menu">
                <Dirs/>
            </div>
            <div className="main">
                right side
                <div>
                    {/*<ComponentOne/>*/}
                    {/*<ComponentTwo/>*/}
                    {/*<ZustandComp/>*/}
                </div>
            </div>
            <footer>
                foot
            </footer>
        </div>
    );
}


export default App;
