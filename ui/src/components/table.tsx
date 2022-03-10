import { ReactNode, useState } from 'react'

interface TableProps {
	title: string
	headers: string[]
	items: ReactNode[] | undefined
}

export function Table({ title, headers, items = [] }: TableProps) {
	return (
		<div>
			<h2>{title}</h2>
			{items?.length === 0 && <div css={{ marginTop: 6 }}>No items</div>}
			<div
				css={{
					padding: 8,
					margin: 12,
					marginBottom: 16,
					display: 'grid',
					gridTemplateColumns: `repeat(${headers.length}, 1fr)`,
					borderBottom: '1px solid #999999',
				}}
			>
				{headers.map((header) => (
					<div key={header}>{header}</div>
				))}
			</div>
			{items}
		</div>
	)
}

interface ItemProps {
	children: ReactNode
	values: string[]
}

export function Item({ children, values }: ItemProps) {
	const [isOpen, setIsOpen] = useState(false)

	return (
		<div
			css={{ backgroundColor: '#eeeeee33', borderRadius: 4, margin: 14, overflow: 'hidden' }}
		>
			<div
				css={{
					backgroundColor: '#eeeeee66',
					padding: 8,
					display: 'grid',
					gridTemplateColumns: `repeat(${values.length}, 1fr)`,
					cursor: 'pointer',
					':hover': {
						backgroundColor: '#eeeeee99',
					},
				}}
				onClick={() => setIsOpen((isOpen) => !isOpen)}
			>
				{values.map((value, index) => (
					<div key={index}>{value}</div>
				))}
			</div>
			{isOpen && <div css={{ padding: 4 }}>{children}</div>}
		</div>
	)
}

interface DetailProps {
	label: string
	value: string
}

export function Detail({ label, value }: DetailProps) {
	if (!value) return null

	return (
		<div css={{ padding: 8 }}>
			<div>
				<span css={{ fontWeight: 500 }}>{label}:</span> {value}
			</div>
		</div>
	)
}
