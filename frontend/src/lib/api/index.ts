import { readFileSync } from "fs";

export interface Config {
  init: string
  frontend: string
  wordpress: string
  server: string
}

export interface Post{
  slug:string
  title:string
  cdate:string
  content:string
  image:string
  views:number
}
export interface Comments{
  name:string
  content:string
  photo:string
  post:string
  cdate:string
}

export interface Index{
  highlight:Post[]
  popular:Post[]
  latest:Post[]
}
export interface SpPost{
  post:Post
  comment:Comment[]
}

export const conf = JSON.parse(String(readFileSync("../config.json"))) as Config
export const getIndex = async () =>{
  const link = conf.server+"/api/post"
  const res = await fetch(link,{
    headers:{
      'Content-Type': 'application/json'
    }
  })
  return await res.json() as Index
}
export const getPost = async (slug:string)=>{
  const link = conf.server+"/api/post"+slug
  const res = await fetch(link,{
    headers:{
      'Content-Type': 'application/json'
    }
  })
  return await res.json() as SpPost
}
