import { defineStore } from 'pinia';

// Import axios to make HTTP requests
import axios from 'axios';
import apiConfig from "../config/api"
import path from "@/lib/path.js";

export const useFilesStore = defineStore('files', {
  state: () => ({
    files: [],
    path: "/",
    loading: false,
    error: '',
  }),
  getters: {
    getFiles(state) {
      return state.files;
    },
    isRoot(state){
      return state.path === "/";
    }
  },
  actions: {
    async fetchFiles(target) {
      this.loading = true;
      try {
        const url  =apiConfig.API_LOCATION+target
        const data = await axios.get(url);
        this.files = data.data.Items;
        this.path = target
      } catch (error) {
        alert(error);
        console.log(error);
      }finally {
        this.loading = false;
      }
    },
    async Navigate(p){
      if (this.loading === false){
        await this.fetchFiles(path.join(this.path,p))
      }

    }
  },
});
