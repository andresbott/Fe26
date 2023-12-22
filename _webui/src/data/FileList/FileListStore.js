import {create} from 'zustand';
import {FileList} from './FileList';
import {join} from "path-browserify";


const _ApiUrl_ = "http://localhost:8080/api/v0/fe"

const usePath = create((set, get) => ({
        loadPath: async (path) => {
            // todo, delegate the action to an external async cuntion, and make sure that data is set in one
            // operation
            try {
                set({loading: true})
                const response = await fetch(_ApiUrl_ + "?path=" + get().path);


                const json = await response.json();

                const fe = new FileList(json)

                set({data: fe})
                // await new Promise(r => setTimeout(r, 1000));
            } catch (error) {
                set({error: error})
            } finally {
                set({loading: false})
            }
        },


        navigate: (dir) => {
            const newPath = join(get().path, dir)
            set({path: newPath})
            get().loadPath(newPath).then(() => {
                // todo here set all the values that changed at once
            })
        },

        navUp: () => {
            const newPath = join(get().path, "../")
            set({path: newPath})
            get().loadPath(newPath)
        },

        // todo add a timestamp with the last loaded and only reload if some time has passed
        path: "",
        loading: true,
        error: null,
        data: {},


        // updateUser: (user) => {
        //     set({data: {...get().data, user: {...get().data?.user, ...user}}})
        // },
        // removeCard: (cardId) => {
        //     const cards = get().data.cards.filter(card => card.id !== cardId);
        //     set({data: {...get().data, cards}})
        // }
    })
)


export {usePath}