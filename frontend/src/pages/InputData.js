import { Typography, CssBaseline, Container, Toolbar, AppBar, Grid, TextField, makeStyles, Button } from "@material-ui/core";
import React, {useState} from "react";
import { useHistory } from "react-router";
import axios from "axios"
import 'bootstrap-css-only/css/bootstrap.min.css'; 


const useStyles = makeStyles((theme)=>({
    paper: {
        
        paddingTop:"150px",
        display: 'flex',
        flexDirection: 'column',
        
      },
      form: {
        width: '100%',
        marginTop: theme.spacing(1),
      },
      submit: {
        width:"100%",
        margin: theme.spacing(3, 0, 2),
      },

}));
function between(x, min, max) {
    return x >= min && x <= max;
  }

function Ztest(){
    const classes = useStyles();
    const history = useHistory();
    const [data, setData] = useState({
        name:"",
        rating: "",
        cuisine:"",
        address:"",
        openingTime:"",
        closingTime:"",
        costForTwo: ""
    });

    function handleChange(event){
        event.preventDefault();
        setData((prev)=>{
            return {...prev, [event.target.name]:event.target.value};
        })
    }
    function handleSubmit(event){
        event.preventDefault();
        console.log(data);
        if(data.name===""||data.rating===""||data.cuisine===""||data.address===""||data.openingTime===""||data.closingTime===""||data.costForTwo===""){
            alert("At least one of the fields has invalid input!");
        }
        else{
            let floatNumber = parseFloat(data.rating);
            if(between(floatNumber, 0, 5)){
                
                let opening=parseInt(data.openingTime);
                let closing=parseInt(data.closingTime);
                if(between(opening, 0, 2400)&&between(closing, opening, 2400)){
                axios.post("http://localhost:5000/add", data)
                .then((response)=>{
                    console.log(response);
                })
                .catch((err)=>{
                    console.log(err);
                })
                .then(()=>{
                    setData({
                        name:"",
                        rating:"",
                        cuisine:"",
                        address:"",
                        openingTime:"",
                        costForTwo:""
                    });
                    history.push("/");
                    
                });
            }
            else{
                alert("Time input invalid!")
            }
        }
        else{
            alert("Rating must be between 0 and 5!")
        }
        }
        
    }

    return<div>
        
            <AppBar position="fixed" display="block" style={{backgroundColor:"black"}}>
                <Toolbar>
                    
                    <Typography variant="h6" className={classes.title} >
                    Submit Details
                    </Typography>
                </Toolbar>
                <nav style={{display:"flex", justifyContent:"center"}}>
                <Button variant="contained" color="primary" href="http://localhost:3000/">
  Go Back
</Button>
                </nav>
            </AppBar>
        
        <div >
        <Container component="main" maxWidth="lg" style={{ alignContent:"center",}}>
            <CssBaseline/>
            <div className={classes.paper} >
                  <Typography variant="h5" align="left">
                    Enter Restaurant Details
                  </Typography>
                  <form className={classes.form} noValidate>
                    <TextField
                        
                        name="name"
                        id="name"
                        label="Name"
                        placeholder="Restaurant name"
                        fullWidth
                        onChange={handleChange}
                        value={data.name}
                        
                    />
                    <TextField
                        
                        name="address"
                        id="address"
                        label="Full Address"
                        placeholder="Full address of the restaurant"
                        fullWidth
                        onChange={handleChange}
                        value={data.address}
                        style={{marginTop:"10px"}}
                        
                    />
                    <TextField
                        
                        name="cuisine"
                        id="cuisine"
                        label="Cuisine"
                        placeholder="Cuisine served by this restaurant"
                        fullWidth
                        onChange={handleChange}
                        value={data.cuisine}
                        style={{marginTop:"10px"}}
                        
                    />
                    <TextField
                        
                        name="rating"
                        id="rating"
                        label="Rating (0 to 5)"
                        placeholder="Restaurant rating"
                        fullWidth
                        onChange={handleChange}
                        value={data.rating}
                        style={{marginTop:"10px"}}
                        
                    />
                    
                    <TextField
                        
                        name="openingTime"
                        id="openingTime"
                        label="Opening Time(24 hour format)"
                        placeholder="Restaurant opening timings"
                        halfWidth
                        onChange={handleChange}
                        value={data.openingTime}
                        style={{marginTop:"10px"}}
                        
                    />
                    <TextField
                        
                        name="closingTime"
                        id="closingTime"
                        label="Closing Time"
                        placeholder="Restaurant closing timings"
                        halfWidth
                        onChange={handleChange}
                        value={data.closingTime}
                        style={{marginTop:"10px", marginLeft: "20px"}}
                        
                    />
                    
                    <TextField
                        
                        name="costForTwo"
                        id="costForTwo"
                        label="CFT in Rupees"
                        placeholder="Cost for two at the restaurant (approx.)"
                        fullWidth
                        onChange={handleChange}
                        value={data.costForTwo}
                        style={{marginTop:"10px"}}
                        
                    />
                    
                    <Button
                        type="submit"
                        halfWidth
                        variant="contained"
                        color="success"
                        className={classes.submit}
                        onClick={handleSubmit}
                        
                        
                        >
                        DONE
                    </Button>
                  </form>  
                </div>    
            

            
                
        </Container>
        </div>
    </div>

}
export default Ztest;