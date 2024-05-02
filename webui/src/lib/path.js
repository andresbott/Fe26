
const path = {
    join: function (a,b){
        if (b===".."){
            let n = a.replace(/\/+$/, ''); //$ marks the end of a string
            n = n.substring(0, n.lastIndexOf("/"))+"/";
            return n
        }
        return a+b+"/"
    }
}

export default path