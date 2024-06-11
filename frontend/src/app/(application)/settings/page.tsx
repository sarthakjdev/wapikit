import { Button } from '~/components/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '~/components/ui/card'
import { ScrollArea } from '~/components/ui/scroll-area'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '~/components/ui/tabs'
import { Input } from '~/components/ui/input'
import { SaveIcon } from 'lucide-react'

export default function SettingsPage() {
	return (
		<ScrollArea className="h-full ">
			<div className="flex-1 space-y-4 p-4 pt-6 md:p-8">
				<div className="flex items-center justify-between space-y-2">
					<h2 className="text-3xl font-bold tracking-tight">Settings</h2>
					<div className="hidden items-center space-x-2 md:flex">
						<Button className="flex flex-row items-center gap-2">
							<SaveIcon className="size-5" />
							Save
						</Button>
					</div>
				</div>
				<Tabs defaultValue="app-settings" className="space-y-4">
					<TabsList>
						<TabsTrigger value="app-settings">App Settings</TabsTrigger>
						<TabsTrigger value="whatsapp-business-account">
							WhatsApp Business Account
						</TabsTrigger>
						<TabsTrigger value="live-chat-settings">Live Chat Settings </TabsTrigger>
						<TabsTrigger value="quick-actions">Quick Actions</TabsTrigger>
						<TabsTrigger value="api-keys">API Keys</TabsTrigger>
					</TabsList>
					<TabsContent value="app-settings" className="space-y-4">
						<Card>
							<CardHeader>
								<CardTitle>Application Name</CardTitle>
								<CardDescription>
									Used to identify your project in the dashboard.
								</CardDescription>
							</CardHeader>
							<CardContent>
								<form>
									<Input placeholder="Project Name" />
								</form>
							</CardContent>
						</Card>
						<Card className="flex flex-row">
							<div className="flex-1">
								<CardHeader>
									<CardTitle>Root Url</CardTitle>
									<CardDescription>
										Used to identify your project in the dashboard.
									</CardDescription>
								</CardHeader>
								<CardContent>
									<form>
										<Input placeholder="Project Name" />
									</form>
								</CardContent>
							</div>
							<div className="tremor-Divider-root mx-auto my-6 flex items-center justify-between gap-3 text-tremor-default text-tremor-content dark:text-dark-tremor-content">
								<div className="h-full w-[1px] bg-tremor-border dark:bg-dark-tremor-border"></div>
							</div>
							<div className="flex-1">
								<CardHeader>
									<CardTitle>Favicon Url </CardTitle>
									<CardDescription>
										Used to identify your project in the dashboard.
									</CardDescription>
								</CardHeader>
								<CardContent>
									<form>
										<Input placeholder="Project Name" />
									</form>
								</CardContent>
							</div>
						</Card>
						<Card className="flex flex-row">
							<div className="flex-1">
								<CardHeader>
									<CardTitle>Media Upload Path</CardTitle>
									<CardDescription>
										Used to identify your project in the dashboard.
									</CardDescription>
								</CardHeader>
								<CardContent>
									<form>
										<Input placeholder="Project Name" />
									</form>
								</CardContent>
							</div>
							<div className="tremor-Divider-root mx-auto my-6 flex items-center justify-between gap-3 text-tremor-default text-tremor-content dark:text-dark-tremor-content">
								<div className="h-full w-[1px] bg-tremor-border dark:bg-dark-tremor-border"></div>
							</div>
							<div className="flex-1">
								<CardHeader>
									<CardTitle>Media Upload URI</CardTitle>
									<CardDescription>
										Used to identify your project in the dashboard.
									</CardDescription>
								</CardHeader>
								<CardContent>
									<form>
										<Input placeholder="Project Name" />
									</form>
								</CardContent>
							</div>
						</Card>
					</TabsContent>
				</Tabs>
			</div>
		</ScrollArea>
	)
}
