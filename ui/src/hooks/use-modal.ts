import { atom, useAtom } from 'jotai'

export enum Modals {
	NodeSettings = 'node-settings',
	EdgeSettings = 'edge-settings',
	TriggerSettings = 'trigger-settings',
	SavePipeline = 'save-pipeline',
	TaskLog = 'task-log',
	NewIntegration = 'new-integration',
	NewTrigger = 'new-trigger',
}

export const modalAtom = atom<{ isOpen: boolean; kind: Modals | null; data: unknown | null }>({
	isOpen: false,
	kind: null,
	data: null,
})

export function useModal() {
	const [modal, setModal] = useAtom(modalAtom)

	return {
		...modal,
		open: (kind: Modals, data: unknown = null) => setModal({ isOpen: true, kind, data }),
		close: () => setModal({ isOpen: false, kind: null, data: null }),
	}
}
