import Tkinter as tk
master = tk.Tk()
# master.geometry('540x300')
tk.Label(master, text="User says").grid(row=0, column=0, sticky=tk.W)
tk.Label(master, text="Olly should say").grid(row=1, column=0, sticky=tk.W)
e1 = tk.Entry(master)
e2 = tk.Entry(master)
e1.grid(row=0, column=1,sticky=tk.W+tk.E)
e2.grid(row=1, column=1, sticky=tk.W+tk.E)
tk.Button(master, text='Quit', command=master.quit).grid(row=3, column=0)
T = tk.Text(master,height = 10, width=60)
T.grid(row=2, columnspan=3)
T.insert(tk.END, "Just a text Widget\nin two lines\n")
T.configure(state='disabled')
tk.mainloop( )