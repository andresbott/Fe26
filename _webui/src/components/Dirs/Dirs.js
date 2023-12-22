import React, {memo} from 'react';
import {usePath} from "../../data/FileList/FileListStore";
import './Dirs.scss';

const Dirs = () => {

    const [data, loading] = usePath(state => [state.data, state.loading])
    if (loading) {
        return <h2>loading...</h2>
    }
    return (
        <MemoizedDirsContent files={data.Dirs} loading={loading}></MemoizedDirsContent>
    );
};

export default Dirs;


const DirList = ({files}) => {
    return (
        <>
            <h2>Directories</h2>
            <DirEntry key={".."} name={"../"}/>
            {files.map((item, index) => (
                <DirEntry key={index} name={item}/>
            ))}
        </>
    );
}

function DirEntry(props) {
    const navigate = usePath(state => state.navigate)
    const handleClick = function () {
        console.log(props.name)
        navigate(props.name)
    }

    return <li className={"DirEntry"}><a onClick={handleClick}>{props.name}</a></li>;
}

const MemoizedDirsContent = memo(
    DirList,
    (prevProps, newProps) => {
        // console.log(newProps.loading)
        return newProps.loading !== false
    } //condition to determine when you want to update component
);
