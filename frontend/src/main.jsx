import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import NavBar from './Components/NavBar/NavBar'
import { RouterProvider, createBrowserRouter, Outlet } from 'react-router-dom'
import App from './Components/App/App'
import Signup from './Components/User/Signup/Signup'
import HomePage from './Components/Pages/HomePage'
import AddArticlePage from './Components/Pages/AddArticlePage'
import ProtectedRoute from './Components/ProtectedRoute/ProtectedRoute'
import ArticleWithId from './Components/Article/ArticleWithId'
import Login from './Components/User/Login/LogIn'


const RootLayout = () => (
  <>
    <NavBar />
    <Outlet /> {/* Ceci rendra le composant spécifique à chaque route */}
  </>
)

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,  // Utilise RootLayout pour afficher NavBar sur toutes les pages
    children: [
      { path: "/", element: <HomePage /> },
      { path: "/:articleId", element: <ArticleWithId /> },
      { path: "/login", element: <Login /> },
      { path: "/signup", element: <Signup /> },
      { path: "/addArticle", element: <AddArticlePage /> },
      /*{ path: "/add-article", element:(
        <ProtectedRoute role="admin">
          <AddArticlePage />  
        </ProtectedRoute>
        )
      }*/
    ]
  }
])

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>
)
