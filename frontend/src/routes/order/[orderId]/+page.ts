import { Load } from "@sveltejs/kit";

export const load: Load = ({fetch, params}) => {
  console.log(params)
}