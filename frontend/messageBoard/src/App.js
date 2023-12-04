
import { DataProvider } from './DataContext';
import Home from './components/Home';
import ErrorPage from './components/ErrorPage';


function App() {
  return (
    <DataProvider>
    <Home />
    <ErrorPage />
  </DataProvider>
  );
}

export default App;
