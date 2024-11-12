import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import Index from './Components/Index/Index'
import NavBar from './Components/NavBar/NavBar'
import { RouterProvider, createBrowserRouter, Outlet } from 'react-router-dom'
import Signup from './Components/User/Signup/Signup'

const RootLayout = () => (
  <>
    <NavBar />
    <Outlet /> {/* Ceci rendra le composant spécifique à chaque route */}
  </>
)

const router = createBrowserRouter([
  {
    path: "/",
    element: <RootLayout />,  // Utilise RootLayout pour afficher NavBar sur toutes les pages
    children: [
      { path: "/", element: <Index /> },
      { path: "/signup", element: <Signup /> }
    ]
  }
])

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>
)
