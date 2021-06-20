import React, { useEffect, useState } from "react";
import {useHistory, usehistory} from "react-router-dom";
import Container from "@material-ui/core/Container";
import { AppBar, CssBaseline, Grid, makeStyles, Typography, Fab, Toolbar } from "@material-ui/core";
import {Add, Menu} from "@material-ui/icons";
import axios from "axios"
import MyCard from "../components/DisplayRestaurant"
import {Accordion, Card} from 'react-bootstrap';
import Button from "@material-ui/core/Button";
const useStyles = makeStyles((theme)=>({
      fab: {
        position: 'fixed',
        bottom: theme.spacing(2),
        right: theme.spacing(2),
      },
      
    
}));


function RestaurantList(){
    const classes = useStyles();
    const history = useHistory();
    const [list, setList] = useState([]);

    useEffect(()=>{
        //get restaurant list when the page loads
        axios.get("http://localhost:5000/restaurants")
        .then(async (response)=>{
            console.log(response.data);
            await setList((prev)=>{
                return response.data;
            });
            console.log(list);
        })
        .catch((err)=>{
            console.log(err);
        });
    },[] );

    return <div>
         <header>
           <div className={classes.bg}>
    <div id='intro-example'
        className='p-5 text-center bg-image'
        style={{ backgroundImage: "url(/food.jpg)",}}
      >
        <div className='mask' style={{ backgroundColor: 'rgba(0, 0, 0, 0.6)' }}>
          <div className='d-flex justify-content-center align-items-center h-100'>
            <div className='text-white'>
              <h1 className='mb-3'>Welcome to Mini-Zomato!</h1>
              <h4 className='mb-3'>Showing Restaurants Near You</h4>
              <nav style={{display:"flex", justifyContent:"center"}}>
                <Button variant="contained" color="primary" href="http://localhost:3000/inputDetails/">
 Add Restaurant
</Button>
                </nav>
            </div>
          </div>
        </div>
      </div>
      </div>
  </header>
       <div className="hello" style={{backgroundColor: 'black', height: "100%"}}>     
        <Container component="main" maxWidth="sm" style={{paddingTop:"30px", backgroundColor: 'black'}}>
            <CssBaseline/>
            <Grid container style={{justifyContent:"center"}}>
                {/* Render all cards */}
                {list.map((item)=>{return <MyCard info = {item}/>})}

            </Grid>
            
                
        </Container>
        </div>
    </div>

}
export default RestaurantList;