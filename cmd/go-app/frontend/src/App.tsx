import { useEffect, useState } from 'react';
import { DomainExpansion, GetData } from "../wailsjs/go/controllers/AllNoteController";
import './App.css';

function App() {
    const [resultText, setResultText] = useState("Please enter your tecnique");
    const [name, setName] = useState('');
    const [isLoading, setIsLoading] = useState<boolean>(false)
    const updateName = (e: any) => setName(e.target.value);
    const updateResultText = (result: string) => setResultText(result);

    function greet() {
        DomainExpansion(name).then(updateResultText);
    }

    async function bye() {
        setIsLoading(true)
        const response = await GetData({ page : 1, pageSize: 10 })
        setResultText(JSON.stringify(response.Data[9].Note_Plain))
        setIsLoading(false)
    }

    useEffect(()=> {
        bye()
    },[])

    return (
        <div >
            <h1>Notas</h1>
            {
                isLoading ? "Cargando..."
                : <div id="result"  dangerouslySetInnerHTML={{ __html: resultText}} />
            }
            
        </div>
    )
}

export default App
