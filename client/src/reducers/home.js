import {PET_ACTIVE,PETS,PETS_LOADING} from '../constants/actions.js'

//pets
//export function pets(state=[],action){
export function pets(state=[{
    id:"123",
    name:"Alexander",
    color:"black and white",
    sex: "male",
    breed:"tabby",
    species: "feline",
    imageURL: "https://s3-us-west-2.amazonaws.com/terraform-in-action/cat-0.jpg"
},{
id:"123123412",
name:"Pike",
color:"black and white",
sex: "male",
breed:"tabby",
species: "feline",
imageURL: "https://s3-us-west-2.amazonaws.com/terraform-in-action/cat-1.jpg"
}],action){
    switch(action.type){
        case PETS:
            return action.payload;
        default:
            return state;
    }
}
export function petsLoading(state=false,action){
    switch(action.type){
        case PETS_LOADING:
        return action.payload;
        default:
        return state;
    }
}
export function petActive(state={},action){
    switch(action.type){
        case PET_ACTIVE:
        return action.payload;
        default:
        return state;
    }
}