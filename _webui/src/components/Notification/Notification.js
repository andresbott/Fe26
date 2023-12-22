import {usePath} from "../../data/FileList/FileListStore";
import React from "react";

const Notification = () => {


    const error = usePath(state => state.error)
    if (error !== null) {
        // alert(error.message)
        console.log("[NOTIFCATION COMPONENT] ERROR: " + error.message)
    }
    return (
        <></>
    );
};

export default Notification;
