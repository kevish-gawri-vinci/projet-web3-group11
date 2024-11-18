
import React from 'react';
import ArticleList from '../Article/ArticleList';
import { useOutletContext } from "react-router-dom";

const HomePage = () => {

    const {articles} = useOutletContext();
    
    return (
        <div>
            <h1>Home Page</h1>
            <div>
                <ArticleList articles={articles}/>
            </div>
        </div>
    );
}
export default HomePage;

