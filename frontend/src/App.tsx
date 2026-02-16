import { useState } from "react"
import "./App.css"

import FeatureTextBox from "./components/FeatureTextBox"
import PredictionField from "./components/PredictionField"
import StackSelector from "./components/StackSelector"
import TeamSizeTextBox from "./components/TeamSizeTextBox"
import TeamSkillSelector from "./components/TeamSkillSelector"

interface Response {
	data: string,
}

export default function App() {
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

		const text: Response = await resp.json()
		setPrediction(text.data)
		setLoading(false)
	}

	return (
		<main className="container">
			<header style={{ gridColumn: '1 / -1' }}>
				<h1>Project Estimator<span style={{ color: 'var(--accent-color)' }}>.</span></h1>
				<p style={{ color: 'var(--text-secondary)', marginTop: '-10px' }}>
					Put your team and predict effort.
				</p>
			</header>

			<div className="left">
				<div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '12px' }}>
					<TeamSizeTextBox value={teamSize} onChange={setTeamSize} />
					<TeamSkillSelector value={seniority} onChange={setSeniority} />
				</div>

				<StackSelector value={stack} onChange={setStack} />
				<FeatureTextBox value={feature} onChange={setFeature} />

				<button onClick={handlePredict} disabled={loading}>
					{loading ? "Calculating..." : "Make Prediction"}
				</button>
			</div>

			<div className="right">
				<PredictionField value={prediction} loading={loading} />
			</div>
		</main>
	)
}