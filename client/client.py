#!/usr/bin/python
# Author: sheppard(ysf1026@gmail.com) 2013-11-21

from Tkinter import *

class App(Frame):
    def about(self):
        w = Label(self.root, text='claw-cgp\nAuthor: sheppard(ysf1026@gmail.com)')
        w.pack(side=TOP)

    def createWidgets(self):
        self.menubar = Menu(self)

        help_menu = Menu(self.menubar, tearoff=0)
        help_menu.add_command(label='About', command=self.about)
        self.menubar.add_cascade(label='Help', menu=help_menu)

        self.menubar.add_command(label='Quit', command=self.quit)

        self.root.config(menu=self.menubar)

    def __init__(self, master=None):
        self.root = master
        master.title('claw-cgp')
        master.geometry('800x600')
        Frame.__init__(self, master)
        self.pack()
        self.createWidgets()

root = Tk()
app = App(master=root)
app.mainloop()
