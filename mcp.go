package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
)

type (
	command      struct{}
	register     struct{}
	memory       struct{}
	debug        struct{}
	assembler    struct{}
	stack        struct{}
	flag         struct{}
	pattern      struct{}
	misc         struct{}
	module       struct{}
	disassembler struct{}

	x64dbg struct {
		Command      command
		Register     register
		Memory       memory
		Debug        debug
		Assembler    assembler
		Stack        stack
		Flag         flag
		Pattern      pattern
		Misc         misc
		Module       module
		Disassembler disassembler
	}
)

/*
static void registercommands()
{
    cmdinit();

    //general purpose
    dbgcmdnew("inc", cbInstrInc, false);
    dbgcmdnew("dec", cbInstrDec, false);
    dbgcmdnew("add", cbInstrAdd, false);
    dbgcmdnew("sub", cbInstrSub, false);
    dbgcmdnew("mul", cbInstrMul, false);
    dbgcmdnew("mulhi", cbInstrMulhi, false);
    dbgcmdnew("div", cbInstrDiv, false);
    dbgcmdnew("and", cbInstrAnd, false);
    dbgcmdnew("or", cbInstrOr, false);
    dbgcmdnew("xor", cbInstrXor, false);
    dbgcmdnew("neg", cbInstrNeg, false);
    dbgcmdnew("not", cbInstrNot, false);
    dbgcmdnew("bswap", cbInstrBswap, false);
    dbgcmdnew("rol", cbInstrRol, false);
    dbgcmdnew("ror", cbInstrRor, false);
    dbgcmdnew("shl,sal", cbInstrShl, false);
    dbgcmdnew("shr", cbInstrShr, false);
    dbgcmdnew("sar", cbInstrSar, false);
    dbgcmdnew("push", cbInstrPush, true);
    dbgcmdnew("pop", cbInstrPop, true);
    dbgcmdnew("popcnt", cbInstrPopcnt, false);
    dbgcmdnew("lzcnt", cbInstrLzcnt, false);
    dbgcmdnew("test", cbInstrTest, false);
    dbgcmdnew("cmp", cbInstrCmp, false);
    dbgcmdnew("mov,set", cbInstrMov, false); //mov a variable, arg1:dest,arg2:src

    //general purpose (SSE/AVX/AVX-512)
    dbgcmdnew("movdqu,movups,movupd", cbInstrMovdqu, true); //move from and to XMM register
    dbgcmdnew("vmovups,vmovupd,vmovdqu", cbInstrVmovdqu, true); //move from and to YMM/ZMM register
    if(detectAVX512())
    {
        dbgcmdnew("kmovq", cbInstrKmovq, true); //move qword from and to K0-K7 register
        dbgcmdnew("kmovd", cbInstrKmovd, true); //move dword from and to K0-K7 register
    }

    //debug control
    dbgcmdnew("InitDebug,init,initdbg", cbDebugInit, false); //init debugger arg1:exefile,[arg2:commandline]
    dbgcmdnew("StopDebug,stop,dbgstop", cbDebugStop, true); //stop debugger
    dbgcmdnew("AttachDebugger,attach", cbDebugAttach, false); //attach
    dbgcmdnew("DetachDebugger,detach", cbDebugDetach, true); //detach
    dbgcmdnew("run,go,r,g", cbDebugRun, true); //unlock WAITID_RUN
    dbgcmdnew("erun,egun,er,eg", cbDebugErun, true); //run + skip first chance exceptions
    dbgcmdnew("serun,sego", cbDebugSerun, true); //run + swallow exception
    dbgcmdnew("pause", cbDebugPause, false); //pause debugger
    dbgcmdnew("DebugContinue,con", cbDebugContinue, true); //set continue status
    dbgcmdnew("StepInto,sti,SingleStep,sstep,sst", cbDebugStepInto, true); //StepInto
    dbgcmdnew("eStepInto,esti", cbDebugeStepInto, true); //StepInto + skip first chance exceptions
    dbgcmdnew("seStepInto,sesti,eSingleStep,esstep,esst", cbDebugseStepInto, true); //StepInto + swallow exception
    dbgcmdnew("StepOver,step,sto,st", cbDebugStepOver, true); //StepOver
    dbgcmdnew("eStepOver,estep,esto,est", cbDebugeStepOver, true); //StepOver + skip first chance exceptions
    dbgcmdnew("seStepOver,sestep,sesto,sest", cbDebugseStepOver, true); //StepOver + swallow exception
    dbgcmdnew("StepOut,rtr", cbDebugStepOut, true); //StepOut
    dbgcmdnew("eStepOut,ertr", cbDebugeStepOut, true); //rtr + skip first chance exceptions
    dbgcmdnew("skip", cbDebugSkip, true); //skip one instruction
    dbgcmdnew("InstrUndo", cbInstrInstrUndo, true); //Instruction undo
    dbgcmdnew("StepUser,StepUserInto", cbDebugStepUserInto, true); // step into until reaching user code
    dbgcmdnew("StepSystem,StepUserInto", cbDebugStepSystemInto, true); // step into until reaching system code

    //breakpoint control
    dbgcmdnew("SetBPX,bp,bpx", cbDebugSetBPX, true); //breakpoint
    dbgcmdnew("DeleteBPX,bpc,bc", cbDebugDeleteBPX, true); //breakpoint delete
    dbgcmdnew("EnableBPX,bpe,be", cbDebugEnableBPX, true); //breakpoint enable
    dbgcmdnew("DisableBPX,bpd,bd", cbDebugDisableBPX, true); //breakpoint disable
    dbgcmdnew("SetHardwareBreakpoint,bph,bphws", cbDebugSetHardwareBreakpoint, true); //hardware breakpoint
    dbgcmdnew("DeleteHardwareBreakpoint,bphc,bphwc", cbDebugDeleteHardwareBreakpoint, true); //delete hardware breakpoint
    dbgcmdnew("EnableHardwareBreakpoint,bphe,bphwe", cbDebugEnableHardwareBreakpoint, true); //enable hardware breakpoint
    dbgcmdnew("DisableHardwareBreakpoint,bphd,bphwd", cbDebugDisableHardwareBreakpoint, true); //disable hardware breakpoint
    dbgcmdnew("SetMemoryBPX,membp,bpm", cbDebugSetMemoryBpx, true); //SetMemoryBPX
    dbgcmdnew("SetMemoryRangeBPX,memrangebp,bpmrange", cbDebugSetMemoryRangeBpx, true); // SetMemoryRangeBpx
    dbgcmdnew("DeleteMemoryBPX,membpc,bpmc", cbDebugDeleteMemoryBreakpoint, true); //delete memory breakpoint
    dbgcmdnew("EnableMemoryBreakpoint,membpe,bpme", cbDebugEnableMemoryBreakpoint, true); //enable memory breakpoint
    dbgcmdnew("DisableMemoryBreakpoint,membpd,bpmd", cbDebugDisableMemoryBreakpoint, true); //enable memory breakpoint
    dbgcmdnew("LibrarianSetBreakpoint,bpdll", cbDebugBpDll, true); //set dll breakpoint
    dbgcmdnew("LibrarianRemoveBreakpoint,bcdll", cbDebugBcDll, true); //remove dll breakpoint
    dbgcmdnew("LibrarianEnableBreakpoint,bpedll", cbDebugBpDllEnable, true); //enable dll breakpoint
    dbgcmdnew("LibrarianDisableBreakpoint,bpddll", cbDebugBpDllDisable, true); //disable dll breakpoint
    dbgcmdnew("SetExceptionBPX", cbDebugSetExceptionBPX, true); //set exception breakpoint
    dbgcmdnew("DeleteExceptionBPX", cbDebugDeleteExceptionBPX, true); //delete exception breakpoint
    dbgcmdnew("EnableExceptionBPX", cbDebugEnableExceptionBPX, true); //enable exception breakpoint
    dbgcmdnew("DisableExceptionBPX", cbDebugDisableExceptionBPX, true); //disable exception breakpoint
    dbgcmdnew("bpgoto", cbDebugSetBPGoto, true);
    dbgcmdnew("bplist", cbDebugBplist, true); //breakpoint list
    dbgcmdnew("SetBPXOptions,bptype", cbDebugSetBPXOptions, false); //breakpoint type

    //conditional breakpoint control
    dbgcmdnew("SetBreakpointName,bpname", cbDebugSetBPXName, true); //set breakpoint name
    dbgcmdnew("SetBreakpointCondition,bpcond,bpcnd", cbDebugSetBPXCondition, true); //set breakpoint breakCondition
    dbgcmdnew("SetBreakpointLog,bplog,bpl", cbDebugSetBPXLog, true); //set breakpoint logText
    dbgcmdnew("SetBreakpointLogCondition,bplogcondition", cbDebugSetBPXLogCondition, true); //set breakpoint logCondition
    dbgcmdnew("SetBreakpointCommand", cbDebugSetBPXCommand, true); //set breakpoint command on hit
    dbgcmdnew("SetBreakpointCommandCondition", cbDebugSetBPXCommandCondition, true); //set breakpoint commandCondition
    dbgcmdnew("SetBreakpointLogFile", cbDebugSetBPXLogFile, true); //set breakpoint logFile
    dbgcmdnew("SetBreakpointFastResume", cbDebugSetBPXFastResume, true); //set breakpoint fast resume
    dbgcmdnew("SetBreakpointSingleshoot", cbDebugSetBPXSingleshoot, true); //set breakpoint singleshoot
    dbgcmdnew("SetBreakpointSilent", cbDebugSetBPXSilent, true); //set breakpoint fast resume
    dbgcmdnew("GetBreakpointHitCount", cbDebugGetBPXHitCount, true); //get breakpoint hit count
    dbgcmdnew("ResetBreakpointHitCount", cbDebugResetBPXHitCount, true); //reset breakpoint hit count

    dbgcmdnew("SetHardwareBreakpointName,bphwname", cbDebugSetBPXHardwareName, true); //set breakpoint name
    dbgcmdnew("SetHardwareBreakpointCondition,bphwcond", cbDebugSetBPXHardwareCondition, true); //set breakpoint breakCondition
    dbgcmdnew("SetHardwareBreakpointLog,bphwlog", cbDebugSetBPXHardwareLog, true); //set breakpoint logText
    dbgcmdnew("SetHardwareBreakpointLogCondition,bphwlogcondition", cbDebugSetBPXHardwareLogCondition, true); //set breakpoint logText
    dbgcmdnew("SetHardwareBreakpointCommand", cbDebugSetBPXHardwareCommand, true); //set breakpoint command on hit
    dbgcmdnew("SetHardwareBreakpointCommandCondition", cbDebugSetBPXHardwareCommandCondition, true); //set breakpoint commandCondition
    dbgcmdnew("SetHardwareBreakpointLogFile", cbDebugSetBPXHardwareLogFile, true); //set breakpoint logFile
    dbgcmdnew("SetHardwareBreakpointFastResume", cbDebugSetBPXHardwareFastResume, true); //set breakpoint fast resume
    dbgcmdnew("SetHardwareBreakpointSingleshoot", cbDebugSetBPXHardwareSingleshoot, true); //set breakpoint singleshoot
    dbgcmdnew("SetHardwareBreakpointSilent", cbDebugSetBPXHardwareSilent, true); //set breakpoint fast resume
    dbgcmdnew("GetHardwareBreakpointHitCount", cbDebugGetBPXHardwareHitCount, true); //get breakpoint hit count
    dbgcmdnew("ResetHardwareBreakpointHitCount", cbDebugResetBPXHardwareHitCount, true); //reset breakpoint hit count

    dbgcmdnew("SetMemoryBreakpointName,bpmname", cbDebugSetBPXMemoryName, true); //set breakpoint name
    dbgcmdnew("SetMemoryBreakpointCondition,bpmcond", cbDebugSetBPXMemoryCondition, true); //set breakpoint breakCondition
    dbgcmdnew("SetMemoryBreakpointLog,bpmlog", cbDebugSetBPXMemoryLog, true); //set breakpoint log
    dbgcmdnew("SetMemoryBreakpointLogCondition,bpmlogcondition", cbDebugSetBPXMemoryLogCondition, true); //set breakpoint logCondition
    dbgcmdnew("SetMemoryBreakpointCommand", cbDebugSetBPXMemoryCommand, true); //set breakpoint command on hit
    dbgcmdnew("SetMemoryBreakpointCommandCondition", cbDebugSetBPXMemoryCommandCondition, true); //set breakpoint commandCondition
    dbgcmdnew("SetMemoryBreakpointLogFile", cbDebugSetBPXMemoryLogFile, true); //set breakpoint logFile
    dbgcmdnew("SetMemoryBreakpointFastResume", cbDebugSetBPXMemoryFastResume, true); //set breakpoint fast resume
    dbgcmdnew("SetMemoryBreakpointSingleshoot", cbDebugSetBPXMemorySingleshoot, true); //set breakpoint singleshoot
    dbgcmdnew("SetMemoryBreakpointSilent", cbDebugSetBPXMemorySilent, true); //set breakpoint fast resume
    dbgcmdnew("GetMemoryBreakpointHitCount", cbDebugGetBPXMemoryHitCount, true); //get breakpoint hit count
    dbgcmdnew("ResetMemoryBreakpointHitCount", cbDebugResetBPXMemoryHitCount, true); //reset breakpoint hit count

    dbgcmdnew("SetLibrarianBreakpointName", cbDebugSetBPXDLLName, true); //set breakpoint name
    dbgcmdnew("SetLibrarianBreakpointCondition", cbDebugSetBPXDLLCondition, true); //set breakpoint breakCondition
    dbgcmdnew("SetLibrarianBreakpointLog", cbDebugSetBPXDLLLog, true); //set breakpoint log
    dbgcmdnew("SetLibrarianBreakpointLogCondition", cbDebugSetBPXDLLLogCondition, true); //set breakpoint logCondition
    dbgcmdnew("SetLibrarianBreakpointCommand", cbDebugSetBPXDLLCommand, true); //set breakpoint command on hit
    dbgcmdnew("SetLibrarianBreakpointCommandCondition", cbDebugSetBPXDLLCommandCondition, true); //set breakpoint commandCondition
    dbgcmdnew("SetLibrarianBreakpointLogFile", cbDebugSetBPXDLLLogFile, true); //set breakpoint logFile
    dbgcmdnew("SetLibrarianBreakpointFastResume", cbDebugSetBPXDLLFastResume, true); //set breakpoint fast resume
    dbgcmdnew("SetLibrarianBreakpointSingleshoot", cbDebugSetBPXDLLSingleshoot, true); //set breakpoint singleshoot
    dbgcmdnew("SetLibrarianBreakpointSilent", cbDebugSetBPXDLLSilent, true); //set breakpoint fast resume
    dbgcmdnew("GetLibrarianBreakpointHitCount", cbDebugGetBPXDLLHitCount, true); //get breakpoint hit count
    dbgcmdnew("ResetLibrarianBreakpointHitCount", cbDebugResetBPXDLLHitCount, true); //reset breakpoint hit count

    dbgcmdnew("SetExceptionBreakpointName", cbDebugSetBPXExceptionName, true); //set breakpoint name
    dbgcmdnew("SetExceptionBreakpointCondition", cbDebugSetBPXExceptionCondition, true); //set breakpoint breakCondition
    dbgcmdnew("SetExceptionBreakpointLog", cbDebugSetBPXExceptionLog, true); //set breakpoint log
    dbgcmdnew("SetExceptionBreakpointLogCondition", cbDebugSetBPXExceptionLogCondition, true); //set breakpoint logCondition
    dbgcmdnew("SetExceptionBreakpointCommand", cbDebugSetBPXExceptionCommand, true); //set breakpoint command on hit
    dbgcmdnew("SetExceptionBreakpointCommandCondition", cbDebugSetBPXExceptionCommandCondition, true); //set breakpoint commandCondition
    dbgcmdnew("SetExceptionBreakpointLogFile", cbDebugSetBPXExceptionLogFile, true); //set breakpoint logFile
    dbgcmdnew("SetExceptionBreakpointFastResume", cbDebugSetBPXExceptionFastResume, true); //set breakpoint fast resume
    dbgcmdnew("SetExceptionBreakpointSingleshoot", cbDebugSetBPXExceptionSingleshoot, true); //set breakpoint singleshoot
    dbgcmdnew("SetExceptionBreakpointSilent", cbDebugSetBPXExceptionSilent, true); //set breakpoint fast resume
    dbgcmdnew("GetExceptionBreakpointHitCount", cbDebugGetBPXExceptionHitCount, true); //get breakpoint hit count
    dbgcmdnew("ResetExceptionBreakpointHitCount", cbDebugResetBPXExceptionHitCount, true); //reset breakpoint hit count

    //tracing
    dbgcmdnew("TraceIntoConditional,ticnd", cbDebugTraceIntoConditional, true); //Trace into conditional
    dbgcmdnew("TraceOverConditional,tocnd", cbDebugTraceOverConditional, true); //Trace over conditional
    dbgcmdnew("TraceIntoBeyondTraceCoverage,TraceIntoBeyondTraceRecord,tibt", cbDebugTraceIntoBeyondTraceRecord, true); //Trace into beyond trace record
    dbgcmdnew("TraceOverBeyondTraceCoverage,TraceOverBeyondTraceRecord,tobt", cbDebugTraceOverBeyondTraceRecord, true); //Trace over beyond trace record
    dbgcmdnew("TraceIntoIntoTraceCoverage,TraceIntoIntoTraceRecord,tiit", cbDebugTraceIntoIntoTraceRecord, true); //Trace into into trace record
    dbgcmdnew("TraceOverIntoTraceCoverage,TraceOverIntoTraceRecord,toit", cbDebugTraceOverIntoTraceRecord, true); //Trace over into trace record
    dbgcmdnew("RunToParty", cbDebugRunToParty, true); //Run to code in a party
    dbgcmdnew("RunToUserCode,rtu", cbDebugRunToUserCode, true); //Run to user code
    dbgcmdnew("TraceSetLog,SetTraceLog", cbDebugTraceSetLog, true); //Set trace log text + condition
    dbgcmdnew("TraceSetCommand,SetTraceCommand", cbDebugTraceSetCommand, true); //Set trace command text + condition
    dbgcmdnew("TraceSetLogFile,SetTraceLogFile", cbDebugTraceSetLogFile, true); //Set trace log file
    dbgcmdnew("StartTraceRecording,StartRunTrace,opentrace", cbDebugStartTraceRecording, true); //start run trace (Ollyscript command "opentrace" "opens run trace window")
    dbgcmdnew("StopTraceRecording,StopRunTrace,tc", cbDebugStopTraceRecording, true); //stop run trace (and Ollyscript command)

    //thread control
    dbgcmdnew("createthread,threadcreate,newthread,threadnew", cbDebugCreatethread, true); //create thread
    dbgcmdnew("switchthread,threadswitch", cbDebugSwitchthread, true); //switch thread
    dbgcmdnew("suspendthread,threadsuspend", cbDebugSuspendthread, true); //suspend thread
    dbgcmdnew("resumethread,threadresume", cbDebugResumethread, true); //resume thread
    dbgcmdnew("killthread,threadkill", cbDebugKillthread, true); //kill thread
    dbgcmdnew("suspendallthreads,threadsuspendall", cbDebugSuspendAllThreads, true); //suspend all threads
    dbgcmdnew("resumeallthreads,threadresumeall", cbDebugResumeAllThreads, true); //resume all threads
    dbgcmdnew("setthreadpriority,setprioritythread,threadsetpriority", cbDebugSetPriority, true); //set thread priority
    dbgcmdnew("threadsetname,setthreadname", cbDebugSetthreadname, true); //set thread name

    //memory operations
    dbgcmdnew("alloc", cbDebugAlloc, true); //allocate memory
    dbgcmdnew("free", cbDebugFree, true); //free memory
    dbgcmdnew("Fill,memset", cbDebugMemset, true); //memset
    dbgcmdnew("memcpy", cbDebugMemcpy, true); //memcpy
    dbgcmdnew("getpagerights,getrightspage", cbDebugGetPageRights, true);
    dbgcmdnew("setpagerights,setrightspage", cbDebugSetPageRights, true);
    dbgcmdnew("savedata", cbInstrSavedata, true); //save data to disk
    dbgcmdnew("minidump", cbInstrMinidump, true); //create a minidump

    //operating system control
    dbgcmdnew("GetPrivilegeState", cbGetPrivilegeState, true); //get priv state
    dbgcmdnew("EnablePrivilege", cbEnablePrivilege, true); //enable priv
    dbgcmdnew("DisablePrivilege", cbDisablePrivilege, true); //disable priv
    dbgcmdnew("handleclose,closehandle", cbHandleClose, true); //close remote handle
    dbgcmdnew("EnableWindow", cbEnableWindow, true); //enable remote window
    dbgcmdnew("DisableWindow", cbDisableWindow, true); //disable remote window

    //watch control
    dbgcmdnew("AddWatch", cbAddWatch, true); // add watch
    dbgcmdnew("DelWatch", cbDelWatch, true); // delete watch
    dbgcmdnew("SetWatchdog", cbSetWatchdog, true); // Setup watchdog
    dbgcmdnew("SetWatchExpression", cbSetWatchExpression, true); // Set watch expression
    dbgcmdnew("SetWatchName", cbSetWatchName, true); // Set watch name
    dbgcmdnew("SetWatchType", cbSetWatchType, true); // Set watch type
    dbgcmdnew("CheckWatchdog", cbCheckWatchdog, true); // Watchdog


    //variables
    dbgcmdnew("varnew,var", cbInstrVar, false); //make a variable arg1:name,[arg2:value]
    dbgcmdnew("vardel", cbInstrVarDel, false); //delete a variable, arg1:variable name
    dbgcmdnew("varlist", cbInstrVarList, false); //list variables[arg1:type filter]

    //searching
    dbgcmdnew("find", cbInstrFind, true); //find a pattern
    dbgcmdnew("findall", cbInstrFindAll, true); //find all patterns
    dbgcmdnew("findallmem,findmemall", cbInstrFindAllMem, true); //memory map pattern find
    dbgcmdnew("findasm,asmfind", cbInstrFindAsm, true); //find instruction
    dbgcmdnew("reffind,findref,ref", cbInstrRefFind, true); //find references to a value
    dbgcmdnew("reffindrange,findrefrange,refrange", cbInstrRefFindRange, true);
    dbgcmdnew("refstr,strref", cbInstrRefStr, true); //find string references
    dbgcmdnew("reffunctionpointer", cbInstrRefFuncionPointer, true); //find function pointers
    dbgcmdnew("modcallfind", cbInstrModCallFind, true); //find intermodular calls
    dbgcmdnew("setmaxfindresult,findsetmaxresult", cbInstrSetMaxFindResult, false); //set the maximum number of occurences found
    dbgcmdnew("guidfind,findguid", cbInstrGUIDFind, true); //find GUID references TODO: undocumented

    //user database
    dbgcmdnew("dbsave,savedb", cbInstrDbsave, true); //save program database
    dbgcmdnew("dbload,loaddb", cbInstrDbload, true); //load program database
    dbgcmdnew("dbclear,cleardb", cbInstrDbclear, true); //clear program database

    dbgcmdnew("commentset,cmt,cmtset", cbInstrCommentSet, true); //set/edit comment
    dbgcmdnew("commentdel,cmtc,cmtdel", cbInstrCommentDel, true); //delete comment
    dbgcmdnew("commentlist", cbInstrCommentList, true); //list comments
    dbgcmdnew("commentclear", cbInstrCommentClear, true); //clear comments

    dbgcmdnew("labelset,lbl,lblset", cbInstrLabelSet, true); //set/edit label
    dbgcmdnew("labeldel,lblc,lbldel", cbInstrLabelDel, true); //delete label
    dbgcmdnew("labellist", cbInstrLabelList, true); //list labels
    dbgcmdnew("labelclear", cbInstrLabelClear, true); //clear labels

    dbgcmdnew("bookmarkset,bookmark", cbInstrBookmarkSet, true); //set bookmark
    dbgcmdnew("bookmarkdel,bookmarkc", cbInstrBookmarkDel, true); //delete bookmark
    dbgcmdnew("bookmarklist", cbInstrBookmarkList, true); //list bookmarks
    dbgcmdnew("bookmarkclear", cbInstrBookmarkClear, true); //clear bookmarks

    dbgcmdnew("functionadd,func", cbInstrFunctionAdd, true); //function
    dbgcmdnew("functiondel,funcc", cbInstrFunctionDel, true); //function
    dbgcmdnew("functionlist", cbInstrFunctionList, true); //list functions
    dbgcmdnew("functionclear", cbInstrFunctionClear, false); //delete all functions

    dbgcmdnew("argumentadd", cbInstrArgumentAdd, true); //add argument
    dbgcmdnew("argumentdel", cbInstrArgumentDel, true); //delete argument
    dbgcmdnew("argumentlist", cbInstrArgumentList, true); //list arguments
    dbgcmdnew("argumentclear", cbInstrArgumentClear, false); //delete all arguments

    dbgcmdnew("loopadd", cbInstrLoopAdd, true); //add loop TODO: undocumented
    dbgcmdnew("loopdel", cbInstrLoopDel, true); //delete loop TODO: undocumented
    dbgcmdnew("looplist", cbInstrLoopList, true); //list loops TODO: undocumented
    dbgcmdnew("loopclear", cbInstrLoopClear, true); //clear loops TODO: undocumented

    //analysis
    dbgcmdnew("analyse,analyze,anal", cbInstrAnalyse, true); //secret analysis command
    dbgcmdnew("exanal,exanalyse,exanalyze", cbInstrExanalyse, true); //exception directory analysis
    dbgcmdnew("cfanal,cfanalyse,cfanalyze", cbInstrCfanalyse, true); //control flow analysis
    dbgcmdnew("analyse_nukem,analyze_nukem,anal_nukem", cbInstrAnalyseNukem, true); //secret analysis command #2
    dbgcmdnew("analxrefs,analx", cbInstrAnalxrefs, true); //analyze xrefs
    dbgcmdnew("analrecur,analr", cbInstrAnalrecur, true); //analyze a single function
    dbgcmdnew("analadv", cbInstrAnalyseadv, true); //analyze xref,function and data
    dbgcmdnew("traceexecute", cbInstrTraceexecute, true); //execute trace record on address TODO: undocumented

    dbgcmdnew("virtualmod", cbInstrVirtualmod, true); //virtual module
    dbgcmdnew("symdownload,downloadsym", cbDebugDownloadSymbol, true); //download symbols
    dbgcmdnew("symload,loadsym", cbDebugLoadSymbol, true); //load symbols
    dbgcmdnew("symunload,unloadsym", cbDebugUnloadSymbol, true); //unload symbols
    dbgcmdnew("imageinfo,modimageinfo", cbInstrImageinfo, true); //print module image information
    dbgcmdnew("GetRelocSize,grs", cbInstrGetRelocSize, true); //get relocation table size
    dbgcmdnew("exhandlers", cbInstrExhandlers, true); //enumerate exception handlers
    dbgcmdnew("exinfo", cbInstrExinfo, true); //dump last exception information

    //types
    dbgcmdnew("DataUnknown", cbInstrDataUnknown, true); //mark as Unknown
    dbgcmdnew("DataByte,db", cbInstrDataByte, true); //mark as Byte
    dbgcmdnew("DataWord,dw", cbInstrDataWord, true); //mark as Word
    dbgcmdnew("DataDword,dd", cbInstrDataDword, true); //mark as Dword
    dbgcmdnew("DataFword", cbInstrDataFword, true); //mark as Fword
    dbgcmdnew("DataQword,dq", cbInstrDataQword, true); //mark as Qword
    dbgcmdnew("DataTbyte", cbInstrDataTbyte, true); //mark as Tbyte
    dbgcmdnew("DataOword", cbInstrDataOword, true); //mark as Oword
    dbgcmdnew("DataMmword", cbInstrDataMmword, true); //mark as Mmword
    dbgcmdnew("DataXmmword", cbInstrDataXmmword, true); //mark as Xmmword
    dbgcmdnew("DataYmmword", cbInstrDataYmmword, true); //mark as Ymmword
    dbgcmdnew("DataFloat,DataReal4,df", cbInstrDataFloat, true); //mark as Float
    dbgcmdnew("DataDouble,DataReal8", cbInstrDataDouble, true); //mark as Double
    dbgcmdnew("DataLongdouble,DataReal10", cbInstrDataLongdouble, true); //mark as Longdouble
    dbgcmdnew("DataAscii,da", cbInstrDataAscii, true); //mark as Ascii
    dbgcmdnew("DataUnicode,du", cbInstrDataUnicode, true); //mark as Unicode
    dbgcmdnew("DataCode,dc", cbInstrDataCode, true); //mark as Code
    dbgcmdnew("DataJunk", cbInstrDataJunk, true); //mark as Junk
    dbgcmdnew("DataMiddle", cbInstrDataMiddle, true); //mark as Middle

    dbgcmdnew("AddType", cbInstrAddType, false); //AddType
    dbgcmdnew("AddStruct", cbInstrAddStruct, false); //AddStruct
    dbgcmdnew("AddUnion", cbInstrAddUnion, false); //AddUnion
    dbgcmdnew("AddMember", cbInstrAddMember, false); //AddMember
    dbgcmdnew("AppendMember", cbInstrAppendMember, false); //AppendMember
    dbgcmdnew("AddFunction", cbInstrAddFunction, false); //AddFunction
    dbgcmdnew("AddArg", cbInstrAddArg, false); //AddArg
    dbgcmdnew("AppendArg", cbInstrAppendArg, false); //AppendArg
    dbgcmdnew("SizeofType", cbInstrSizeofType, false); //SizeofType
    dbgcmdnew("VisitType,DisplayType,dt", cbInstrVisitType, false); //VisitType
    dbgcmdnew("ClearTypes", cbInstrClearTypes, false); //ClearTypes
    dbgcmdnew("RemoveType", cbInstrRemoveType, false); //RemoveType
    dbgcmdnew("EnumTypes", cbInstrEnumTypes, false); //EnumTypes
    dbgcmdnew("LoadTypes", cbInstrLoadTypes, false); //LoadTypes
    dbgcmdnew("ParseTypes", cbInstrParseTypes, false); //ParseTypes

    //plugins
    dbgcmdnew("StartScylla,scylla,imprec", cbDebugStartScylla, false); //start scylla
    dbgcmdnew("plugload,pluginload,loadplugin", cbInstrPluginLoad, false); //load plugin
    dbgcmdnew("plugunload,pluginunload,unloadplugin", cbInstrPluginUnload, false); //unload plugin
    dbgcmdnew("plugreload,pluginreload,reloadplugin", cbInstrPluginReload, false); //reload plugin

    //script
    dbgcmdnew("scriptload", cbScriptLoad, false);
    dbgcmdnew("msg", cbScriptMsg, false);
    dbgcmdnew("msgyn", cbScriptMsgyn, false);
    dbgcmdnew("log", cbInstrLog, false); //log command with superawesome hax
    dbgcmdnew("htmllog", cbInstrHtmlLog, false); //command for testing
    dbgcmdnew("scriptdll,dllscript", cbScriptDll, false); //execute a script DLL
    dbgcmdnew("scriptcmd", cbScriptCmd, false); // execute a script command TODO: undocumented

    //gui
    dbgcmdnew("showthreadid", cbShowThreadId, false); // show given thread in threads
    dbgcmdnew("disasm,dis,d", cbDebugDisasm, true); //doDisasm
    dbgcmdnew("dump", cbDebugDump, true); //dump at address
    dbgcmdnew("sdump", cbDebugStackDump, true); //dump at stack address
    dbgcmdnew("memmapdump", cbDebugMemmapdump, true);
    dbgcmdnew("graph", cbInstrGraph, true); //graph function
    dbgcmdnew("guiupdateenable", cbInstrEnableGuiUpdate, true); //enable gui message
    dbgcmdnew("guiupdatedisable", cbInstrDisableGuiUpdate, true); //disable gui message
    dbgcmdnew("setfreezestack", cbDebugSetfreezestack, false); //freeze the stack from auto updates
    dbgcmdnew("refinit", cbInstrRefinit, false);
    dbgcmdnew("refadd", cbInstrRefadd, false);
    dbgcmdnew("refget", cbInstrRefGet, false);
    dbgcmdnew("EnableLog,LogEnable", cbInstrEnableLog, false); //enable log
    dbgcmdnew("DisableLog,LogDisable", cbInstrDisableLog, false); //disable log
    dbgcmdnew("ClearLog,cls,lc,lclr", cbClearLog, false); //clear the log
    dbgcmdnew("SaveLog,LogSave", cbSaveLog, false); //save the log
    dbgcmdnew("RedirectLog,LogRedirect", cbRedirectLog, false); //redirect the log
    dbgcmdnew("StopRedirectLog,LogRedirectStop", cbStopRedirectLog, false); //stop redirecting the log
    dbgcmdnew("AddFavouriteTool", cbInstrAddFavTool, false); //add favourite tool
    dbgcmdnew("AddFavouriteCommand", cbInstrAddFavCmd, false); //add favourite command
    dbgcmdnew("AddFavouriteToolShortcut,SetFavouriteToolShortcut", cbInstrSetFavToolShortcut, false); //set favourite tool shortcut
    dbgcmdnew("FoldDisassembly", cbInstrFoldDisassembly, true); //fold disassembly segment
    dbgcmdnew("guiupdatetitle", cbDebugUpdateTitle, true); // set relevant disassembly title
    dbgcmdnew("showref", cbShowReferences, false); // show references window
    dbgcmdnew("symfollow", cbSymbolsFollow, false); // follow address in symbols tab
    dbgcmdnew("gototrace,tracegoto", cbGotoTrace, false); // goto index in trace tab

    //misc
    dbgcmdnew("chd", cbInstrChd, false); //Change directory
    dbgcmdnew("zzz,doSleep", cbInstrZzz, false); //sleep

    dbgcmdnew("HideDebugger,dbh,hide", cbDebugHide, true); //HideDebugger
    dbgcmdnew("loadlib", cbDebugLoadLib, true); //Load DLL
    dbgcmdnew("freelib", cbDebugFreeLib, true); //Unload DLL TODO: undocumented
    dbgcmdnew("asm", cbInstrAssemble, true); //assemble instruction
    dbgcmdnew("gpa", cbInstrGpa, true); //get proc address

    dbgcmdnew("setjit,jitset", cbDebugSetJIT, false); //set JIT
    dbgcmdnew("getjit,jitget", cbDebugGetJIT, false); //get JIT
    dbgcmdnew("getjitauto,jitgetauto", cbDebugGetJITAuto, false); //get JIT Auto
    dbgcmdnew("setjitauto,jitsetauto", cbDebugSetJITAuto, false); //set JIT Auto

    dbgcmdnew("getcommandline,getcmdline", cbDebugGetCmdline, true); //Get CmdLine
    dbgcmdnew("setcommandline,setcmdline", cbDebugSetCmdline, true); //Set CmdLine

    dbgcmdnew("mnemonichelp", cbInstrMnemonichelp, false); //mnemonic help
    dbgcmdnew("mnemonicbrief", cbInstrMnemonicbrief, false); //mnemonic brief

    dbgcmdnew("config", cbInstrConfig, false); //get or set config uint
    dbgcmdnew("restartadmin,runas,adminrestart", cbInstrRestartadmin, false); //restart x64dbg as administrator

    //undocumented
    dbgcmdnew("bench", cbDebugBenchmark, true); //benchmark test (readmem etc)
    dbgcmdnew("dprintf", cbPrintf, false); //printf
    dbgcmdnew("setstr,strset", cbInstrSetstr, false); //set a string variable
    dbgcmdnew("getstr,strget", cbInstrGetstr, false); //get a string variable
    dbgcmdnew("copystr,strcpy", cbInstrCopystr, true); //write a string variable to memory
    dbgcmdnew("zydis", cbInstrZydis, true); //disassemble using zydis
    dbgcmdnew("visualize", cbInstrVisualize, true); //visualize analysis
    dbgcmdnew("meminfo", cbInstrMeminfo, true); //command to debug memory map bugs
    dbgcmdnew("briefcheck", cbInstrBriefcheck, true); //check if mnemonic briefs are missing
    dbgcmdnew("focusinfo", cbInstrFocusinfo, false);
    dbgcmdnew("printstack,logstack", cbInstrPrintStack, true); //print the call stack
    dbgcmdnew("flushlog", cbInstrFlushlog, false); //flush the log
    dbgcmdnew("AnimateWait", cbInstrAnimateWait, true); //Wait for the debuggee to pause.
    dbgcmdnew("dbdecompress", cbInstrDbdecompress, false); //Decompress a database.
    dbgcmdnew("DebugFlags", cbInstrDebugFlags, false); //Set ntdll LdrpDebugFlags
    dbgcmdnew("LabelRuntimeFunctions", cbInstrLabelRuntimeFunctions, true); //Label exception directory entries
    dbgcmdnew("cmdtest", cbInstrCmdTest, false); //log argv verbatim
};

*/

func (x x64dbg) Restart() { x.Command.Exec("restartadmin") }
func (x x64dbg) FindAsm(address int, instruction string) (size int, data HexBytes) {
	result := x.Command.Exec("findasm " + strconv.Quote(instruction) + "," + fmt.Sprintf("0x%x", address))
	if result == "" {
		return 0, nil
	}
	var r assemblerResult
	json.Unmarshal([]byte(result), &r)
	return r.Size, r.Data
}

func (command) Exec(cmd string) string {
	return request[string]("ExecCommand", map[string]string{"cmd": cmd})
}

func (debug) Active() bool {
	return request[bool]("IsDebugActive", nil)
}

func (debug) Debugging() bool {
	return request[bool]("Is_Debugging", nil)
}

func (memory) Read(address int, size uint) HexBytes {
	return request[HexBytes]("Memory/Read", map[string]string{"addr": fmt.Sprintf("0x%x", address), "size": fmt.Sprintf("%d", size)})
}

func (memory) Write(address int, data HexBytes) bool {
	return request[bool]("Memory/Write", map[string]string{"addr": fmt.Sprintf("0x%x", address), "data": hex.EncodeToString(data)})
}

func (memory) IsValidPtr(address int) bool {
	return request[bool]("Memory/IsValidPtr", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}

func (memory) GetProtectFlag(address int) string { //todo gen enum flag
	return request[string]("Memory/GetProtect", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}

type void any

func (debug) Run()      { request[void]("Debug/Run", nil) }
func (debug) Pause()    { request[void]("Debug/Pause", nil) }
func (debug) Stop()     { request[void]("Debug/Stop", nil) }
func (debug) StepIn()   { request[void]("Debug/StepIn", nil) }
func (debug) StepOver() { request[void]("Debug/StepOver", nil) }
func (debug) StepOut()  { request[void]("Debug/StepOut", nil) }
func (debug) SetBreakpoint(address int) bool { //todo 添加硬件断点
	return request[bool]("Debug/SetBreakpoint", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debug) DeleteBreakpoint(address int) bool {
	return request[bool]("Debug/DeleteBreakpoint", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}

type assemblerResult struct {
	Status string   `json:"success"`
	Size   int      `json:"size"`
	Data   HexBytes `json:"bytes"`
}

func (assembler) Assemble(address int, instruction string) assemblerResult {
	return request[assemblerResult]("Assembler/Assemble", map[string]string{"addr": fmt.Sprintf("0x%x", address), "instruction": instruction})
}
func (assembler) AssembleMem(address int, instructionOpcodes HexBytes) bool {
	return request[bool]("Assembler/AssembleMem", map[string]string{"addr": fmt.Sprintf("0x%x", address), "instruction": hex.EncodeToString(instructionOpcodes)})
}

func (stack) Pop() HexInt { //todo 改成泛型
	return request[HexInt]("Stack/Pop", nil)
}
func (stack) Push(value uint) HexInt {
	return request[HexInt]("Stack/Push", map[string]string{"value": fmt.Sprintf("0x%x", value)})
}
func (stack) Peek(offset int) HexInt {
	return request[HexInt]("Stack/Peek", map[string]string{"offset": fmt.Sprintf("0x%x", offset)})
}

func (disassembler) AtAddress(address int) disassemblerAddress {
	return request[disassemblerAddress]("Disasm/GetInstruction", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (disassembler) AtAddressWithSize(address int, size int) []disassemblerAddress {
	if size < 1 || size > 100 {
		panic("count should be between 1 and 100 bytes buffer")
	}
	return request[[]disassemblerAddress]("Disasm/GetInstructionRange", map[string]string{"addr": fmt.Sprintf("0x%x", address), "count": fmt.Sprintf("%d", size)})
}
func (disassembler) AtRip() disassembleRip {
	return request[disassembleRip]("Disasm/GetInstructionAtRIP", nil)
}
func (disassembler) AtRipFromStepIn() disassembleRipWithSetupIn {
	return request[disassembleRipWithSetupIn]("Disasm/StepInWithDisasm", nil)
}

type disassemblerAddress struct {
	Address     HexInt `json:"address"`
	Instruction string `json:"instruction"`
	Size        int    `json:"size"`
}
type disassembleRip struct {
	Rip         HexInt `json:"rip"`
	Instruction string `json:"instruction"`
	Size        int    `json:"size"`
}
type disassembleRipWithSetupIn struct {
	StepResult  string `json:"step_result"`
	Rip         HexInt `json:"rip"`
	Instruction string `json:"instruction"`
	Size        int    `json:"size"`
}

// Get flag: Flag name (ZF, OF, CF, PF, SF, TF, AF, DF, IF)
// todo gen enum flag
func (flag) Get(name string) bool {
	return request[bool]("Flag/Get", map[string]string{"flag": name})
}

func (flag) Set(name string, value bool) string {
	return request[string]("Flag/Set", map[string]string{"flag": name, "value": fmt.Sprintf("%v", value)})
}

// FindMemory todo 特征码支持字节切片类型
func (pattern) FindMemory(start int, size int, pattern string) (address HexInt) {
	return request[HexInt]("Pattern/FindMem", map[string]string{"start": fmt.Sprintf("0x%x", start), "size": fmt.Sprintf("%d", size), "pattern": pattern})
}

func (misc) ParseExpression(expression string) (value uint) {
	return request[uint]("Misc/ParseExpression", map[string]string{"expression": expression})
}

func (misc) GetApiAddressFromModule(module string, api string) (address HexInt) {
	return request[HexInt]("Misc/RemoteGetProcAddress", map[string]string{"module": module, "api": api})
}

type memoryBase struct {
	BaseAddress HexInt `json:"base_address"`
	Size        uint   `json:"size"`
}

func (memory) FindBaseByAddress(address int) memoryBase {
	return request[memoryBase]("MemoryBase", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}

type moduleInfo struct {
	BaseAddress  HexInt `json:"base_address"`
	Size         uint   `json:"size"`
	Entry        HexInt `json:"entry"`
	SectionCount int    `json:"section_count"`
	Name         string
	Path         string
}

type moduleSectionInfo struct {
	Address HexInt `json:"address"`
	Size    uint   `json:"size"`
	Name    string
}

type moduleExport struct {
	Ordinal         HexInt `json:"ordinal"`
	Rva             HexInt `json:"rva"`
	Va              HexInt `json:"va"`
	Forwarded       bool   `json:"forwarded"`
	ForwardName     string
	Name            string
	UndecoratedName string
}

type moduleImport struct {
	IatRva          HexInt `json:"iat_rva"`
	IatVa           HexInt `json:"iat_va"`
	Ordinal         HexInt `json:"ordinal"`
	Name            string
	UndecoratedName string
}

// todo implement other method in cpp server

func (module) InfoFromAddr(address int) moduleInfo {
	return request[moduleInfo]("Module/InfoFromAddr", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (module) InfoFromName(name string) moduleInfo {
	return request[moduleInfo]("Module/InfoFromName", map[string]string{"name": name})
}
func (module) BaseFromAddr(address int) HexInt {
	return request[HexInt]("Module/BaseFromAddr", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (module) BaseFromName(name string) HexInt {
	return request[HexInt]("Module/BaseFromName", map[string]string{"name": name})
}
func (module) SizeFromAddr(address int) HexInt {
	return request[HexInt]("Module/SizeFromAddr", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (module) SizeFromName(name string) HexInt {
	return request[HexInt]("Module/SizeFromName", map[string]string{"name": name})
}
func (module) NameFromAddr(address int) string {
	return request[string]("Module/NameFromAddr", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (module) PathFromAddr(address int) string {
	return request[string]("Module/PathFromAddr", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (module) PathFromName(name string) string {
	return request[string]("Module/PathFromName", map[string]string{"name": name})
}
func (module) EntryFromAddr(address int) HexInt {
	return request[HexInt]("Module/EntryFromAddr", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (module) EntryFromName(name string) HexInt {
	return request[HexInt]("Module/EntryFromName", map[string]string{"name": name})
}
func (module) SectionCountFromAddr(address int) HexInt {
	return request[HexInt]("Module/SectionCountFromAddr", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (module) SectionCountFromName(name string) HexInt {
	return request[HexInt]("Module/SectionCountFromName", map[string]string{"name": name})
}
func (module) SectionFromAddr(address int, number int) moduleSectionInfo {
	return request[moduleSectionInfo]("Module/SectionFromAddr", map[string]string{"addr": fmt.Sprintf("0x%x", address), "number": fmt.Sprintf("%d", number)})
}
func (module) SectionFromName(name string, number int) moduleSectionInfo {
	return request[moduleSectionInfo]("Module/SectionFromName", map[string]string{"name": name, "number": fmt.Sprintf("%d", number)})
}
func (module) SectionListFromAddr(address int) []moduleSectionInfo {
	return request[[]moduleSectionInfo]("Module/SectionListFromAddr", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (module) SectionListFromName(name string) []moduleSectionInfo {
	return request[[]moduleSectionInfo]("Module/SectionListFromName", map[string]string{"name": name})
}
func (module) GetMainModuleInfo() moduleInfo {
	return request[moduleInfo]("Module/GetMainModuleInfo", nil)
}
func (module) GetMainModuleBase() HexInt {
	return request[HexInt]("Module/GetMainModuleBase", nil)
}
func (module) GetMainModuleSize() HexInt {
	return request[HexInt]("Module/GetMainModuleSize", nil)
}
func (module) GetMainModuleEntry() HexInt {
	return request[HexInt]("Module/GetMainModuleEntry", nil)
}
func (module) GetMainModuleSectionCount() int {
	return request[int]("Module/GetMainModuleSectionCount", nil)
}
func (module) GetMainModuleName() string {
	return request[string]("Module/GetMainModuleName", nil)
}
func (module) GetMainModulePath() string {
	return request[string]("Module/GetMainModulePath", nil)
}
func (module) GetMainModuleSectionList() []moduleSectionInfo {
	return request[[]moduleSectionInfo]("Module/GetMainModuleSectionList", nil)
}
func (module) GetList() []moduleInfo {
	return request[[]moduleInfo]("Module/GetList", nil)
}
func (module) GetExports(mod moduleInfo) []moduleExport {
	return request[[]moduleExport]("Module/GetExports", map[string]string{"mod": fmt.Sprintf("%v", mod)})
}
func (module) GetImports(mod moduleInfo) []moduleImport {
	return request[[]moduleImport]("Module/GetImports", map[string]string{"mod": fmt.Sprintf("%v", mod)})
}
