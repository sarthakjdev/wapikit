'use client'
import { Button } from '~/components/ui/button'
import {
	Form,
	FormControl,
	FormField,
	FormItem,
	FormLabel,
	FormMessage
} from '~/components/ui/form'
import { Input } from '~/components/ui/input'
import {
	Select,
	SelectContent,
	SelectItem,
	SelectTrigger,
	SelectValue
} from '~/components/ui/select'
import { zodResolver } from '@hookform/resolvers/zod'
import { Trash } from 'lucide-react'
import { useRouter } from 'next/navigation'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { type z } from 'zod'
import { useToast } from '../ui/use-toast'
import {
	useGetContactLists,
	useGetOrganizationTags,
	useCreateContact,
	useUpdateContactById,
	ContactStatusEnum,
	type ContactSchema
} from 'root/.generated'
import { Textarea } from '../ui/textarea'
import { NewContactFormSchema } from '~/schema'
import { listStringEnumMembers } from 'ts-enum-utils'

interface FormProps {
	initialData: ContactSchema | null
}

const NewContactForm: React.FC<FormProps> = ({ initialData }) => {
	const router = useRouter()
	const { toast } = useToast()
	const [loading, setLoading] = useState(false)
	const toastMessage = initialData ? 'Product updated.' : 'Product created.'
	const action = initialData ? 'Save changes' : 'Create'

	const listsResponse = useGetContactLists({
		order: 'asc',
		page: 1,
		per_page: 50
	})

	const createContact = useCreateContact()
	const updateContact = useUpdateContactById()

	const defaultValues = initialData
		? {
				...initialData,
				lists: initialData.lists.map(list => list.uniqueId)
			}
		: {
				name: '',
				attributes: '',
				phone: '',
				lists: [],
				status: ContactStatusEnum.Active
			}

	const form = useForm<z.infer<typeof NewContactFormSchema>>({
		resolver: zodResolver(NewContactFormSchema),
		defaultValues
	})

	const onSubmit = async (data: z.infer<typeof NewContactFormSchema>) => {
		try {
			setLoading(true)
			if (initialData) {
				const response = await updateContact.mutateAsync({
					id: initialData.uniqueId,
					data: {
						name: data.name,
						attributes: data.attributes,
						phone: data.phone,
						status: data.status
					}
				})

				if (response.contact.uniqueId) {
					toast({
						variant: 'default',
						title: 'Success!',
						description: toastMessage
					})
				} else {
					toast({
						variant: 'destructive',
						title: 'Uh oh! Something went wrong.',
						description: 'There was a problem with your request.'
					})
				}

				// await axios.post(`/api/products/edit-product/${initialData._id}`, data);
			} else {
				const response = await createContact.mutateAsync(
					{
						data: {
							name: data.name,
							attributes: data.attributes,
							phone: data.phone,
							status: data.status,
							listsIds: data.lists
						}
					},
					{
						onError(error) {
							toast({
								variant: 'destructive',
								title: 'Uh oh! Something went wrong.',
								description: error.message
							})
						}
					}
				)

				if (response.contact.uniqueId) {
					toast({
						variant: 'default',
						title: 'Success!',
						description: toastMessage
					})
				} else {
					toast({
						variant: 'destructive',
						title: 'Uh oh! Something went wrong.',
						description: 'There was a problem with your request.'
					})
				}
			}
			router.refresh()
			router.push(`/dashboard/contacts`)
			toast({
				variant: 'destructive',
				title: 'Uh oh! Something went wrong.',
				description: 'There was a problem with your request.'
			})
		} catch (error: any) {
			toast({
				variant: 'destructive',
				title: 'Uh oh! Something went wrong.',
				description: 'There was a problem with your request.'
			})
		} finally {
			setLoading(false)
		}
	}

	return (
		<>
			<div className="flex flex-1 items-center justify-between">
				{initialData && (
					<Button
						disabled={loading}
						variant="destructive"
						size="sm"
						onClick={() => {
							// ! TODO: headless UI alert modal here
						}}
					>
						<Trash className="h-4 w-4" />
					</Button>
				)}
			</div>
			<Form {...form}>
				<form onSubmit={form.handleSubmit(onSubmit)} className="w-full space-y-8">
					<div className="flex flex-col gap-8">
						<FormField
							control={form.control}
							name="name"
							render={({ field }) => (
								<FormItem>
									<FormLabel>Name</FormLabel>
									<FormControl>
										<Input
											disabled={loading}
											placeholder="Contact title"
											{...field}
										/>
									</FormControl>
									<FormMessage />
								</FormItem>
							)}
						/>
						<FormField
							control={form.control}
							name="description"
							render={({ field }) => (
								<FormItem>
									<FormLabel>Description</FormLabel>
									<FormControl>
										<Textarea
											disabled={loading}
											placeholder="Contact description"
											{...field}
										/>
									</FormControl>
									<FormMessage />
								</FormItem>
							)}
						/>
						<FormField
							control={form.control}
							name="phone"
							render={({ field }) => (
								<FormItem>
									<FormLabel>Phone Number</FormLabel>
									<FormControl>
										<Input
											disabled={loading}
											placeholder="Contact Phone Number"
											{...field}
										/>
									</FormControl>
									<FormMessage />
								</FormItem>
							)}
						/>
						<FormField
							control={form.control}
							name="lists"
							render={({ field }) => (
								<FormItem>
									<FormLabel>Lists</FormLabel>
									<Select
										disabled={loading}
										onValueChange={field.onChange}
										value={field.value.join(',')}
										// defaultValue={field.value}
									>
										<FormControl>
											<SelectTrigger>
												<SelectValue
													defaultValue={field.value}
													placeholder="Select list"
												/>
											</SelectTrigger>
										</FormControl>
										<SelectContent>
											{!listsResponse.data?.lists ||
											listsResponse.data.lists.length === 0 ? (
												<SelectItem value={'no list'} disabled>
													No Lists created yet.
												</SelectItem>
											) : (
												<>
													{listsResponse.data?.lists.map(list => (
														<SelectItem
															key={list.uniqueId}
															value={list.uniqueId}
														>
															{list.name}
														</SelectItem>
													))}
												</>
											)}
										</SelectContent>
									</Select>
									<FormMessage />
								</FormItem>
							)}
						/>
						<FormField
							control={form.control}
							name="status"
							render={({ field }) => (
								<FormItem>
									<FormLabel>Status</FormLabel>
									<Select
										disabled={loading}
										onValueChange={field.onChange}
										value={field.value}
										defaultValue={field.value}
									>
										<FormControl>
											<SelectTrigger>
												<SelectValue
													defaultValue={field.value}
													placeholder="Status"
												/>
											</SelectTrigger>
										</FormControl>
										<SelectContent>
											{listStringEnumMembers(ContactStatusEnum).map(
												status => {
													return (
														<SelectItem
															key={status.name}
															value={status.value}
														>
															{status.name}
														</SelectItem>
													)
												}
											)}
										</SelectContent>
									</Select>
									<FormMessage />
								</FormItem>
							)}
						/>
					</div>
					<Button disabled={loading} className="ml-auto" type="submit">
						{action}
					</Button>
				</form>
			</Form>
		</>
	)
}

export default NewContactForm