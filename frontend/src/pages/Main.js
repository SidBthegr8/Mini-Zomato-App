import React from "react";
import {BrowserRouter, Switch, Route} from "react-router-dom";
import RestaurantListing from "./ShowData";
import Ztest from "./InputData";

function Main(){
    return <div>
        <BrowserRouter>
            <Switch>
                <Route exact path="/">
                    <RestaurantListing />
                </Route>
                <Route exact path="/inputDetails">
                    <Ztest />
                </Route>
        </Switch>
        </BrowserRouter>
        
    </div>
}

export default Main;