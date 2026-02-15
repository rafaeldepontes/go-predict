import { useEffect, useState } from "react"

interface Seniority {
    id: number,
    data: string,
}

interface Props {
    value: string
    onChange: (v: string) => void
}

export default function TeamSkillSelector({ value, onChange }: Props) {
    const [options, setOptions] = useState<Seniority[]>([])

    useEffect(() => {
        const cached = localStorage.getItem("seniorities")
        if (cached) {
            // eslint-disable-next-line react-hooks/set-state-in-effect
            setOptions(JSON.parse(cached))
            return
        }

        fetch(`${import.meta.env.VITE_BACKEND_URL}/seniorities`)
            .then(r => r.json())
            .then(data => {
                localStorage.setItem("seniorities", JSON.stringify(data))
                setOptions(data)
            })
    }, [])

    return (
        <div>
            <p>Seniority:</p>
            <select value={value} onChange={(e) => onChange(e.target.value)}>
                {options.map(o => (
                    <option key={o.id} value={o.data}>{o.data}</option>
                ))}
            </select>
        </div>
    )
}
