import { useState } from "react"
import "./App.css"

import FeatureTextBox from "./components/FeatureTextBox"
import PredictionField from "./components/PredictionField"
import StackSelector from "./components/StackSelector"
import TeamSizeTextBox from "./components/TeamSizeTextBox"
import TeamSkillSelector from "./components/TeamSkillSelector"

function App() {
  const [teamSize, setTeamSize] = useState(1)
  const [seniority, setSeniority] = useState("")
  const [stack, setStack] = useState<string[]>([])
  const [feature, setFeature] = useState("")
  const [prediction, setPrediction] = useState("")
  const [loading, setLoading] = useState(false)

  const handlePredict = async () => {
    setLoading(true)
    setPrediction("")

    let stackReq: string = ""
    stack.forEach((sk) => stackReq += sk + " ")
    stackReq = stackReq.trimEnd()

    const resp = await fetch(`${import.meta.env.VITE_BACKEND_URL}/predict`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        body: feature,
        size: teamSize,
        stack: stackReq,
        level: seniority,
      }),
    })

    const text = await resp.text()
    setPrediction(text)
    setLoading(false)
  }

  return (
    <div className="container">
      <div className="left">
        <div className="row">
          <TeamSizeTextBox value={teamSize} onChange={setTeamSize} />
          <TeamSkillSelector value={seniority} onChange={setSeniority} />
        </div>

        <StackSelector value={stack} onChange={setStack} />
        <FeatureTextBox value={feature} onChange={setFeature} />
        <button onClick={handlePredict}>Make prediction</button>
      </div>

      <div className="right">
        <PredictionField value={prediction} loading={loading} />
      </div>
    </div>
  )
}

export default App
